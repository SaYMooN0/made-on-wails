<template>
  <form @submit.prevent="handleSubmit" class="form-container">
    <div class='crafting-table-letters-zone'>
      <DefCheckBox v-model="isShapelessValue" class="is-shapeless-checkbox" />

      <div class='crafting-table-letters-container'>
        <div v-for="(item, index) in letterItems" :key="index" class='crafting-table-letter-item'
          @dragstart="handleDragStart($event, item.letter)" draggable="true">
          <label class='letter-label'>{{ item.letter }}</label>
          <InputWithSuggestions :value="item.value" @updateValue="item.value = $event" suggestion-type="item"
            class="item-for-letter-input" />
          <div class='letter-delete-button-container' @click="removeLetter(index)">
            <svg class='letter-delete-button' viewBox='0 0 24 24' fill='none'>
              <path d='M20.5001 6H3.5' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
              <path
                d='M18.8332 8.5L18.3732 15.3991C18.1962 18.054 18.1077 19.3815 17.2427 20.1907C16.3777 21 15.0473 21 12.3865 21H11.6132C8.95235 21 7.62195 21 6.75694 20.1907C5.89194 19.3815 5.80344 18.054 5.62644 15.3991L5.1665 8.5'
                stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
              <path d='M9.5 11L10 16' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
              <path d='M14.5 11L14 16' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
              <path
                d='M6.5 6C6.55588 6 6.58382 6 6.60915 5.99936C7.43259 5.97849 8.15902 5.45491 8.43922 4.68032C8.44784 4.65649 8.45667 4.62999 8.47434 4.57697L8.57143 4.28571C8.65431 4.03708 8.69575 3.91276 8.75071 3.8072C8.97001 3.38607 9.37574 3.09364 9.84461 3.01877C9.96213 3 10.0932 3 10.3553 3H13.6447C13.9068 3 14.0379 3 14.1554 3.01877C14.6243 3.09364 15.03 3.38607 15.2493 3.8072C15.3043 3.91276 15.3457 4.03708 15.4286 4.28571L15.5257 4.57697C15.5433 4.62992 15.5522 4.65651 15.5608 4.68032C15.841 5.45491 16.5674 5.97849 17.3909 5.99936C17.4162 6 17.4441 6 17.5 6'
                stroke='#1C274C' stroke-width='1.5' />
            </svg>
          </div>
        </div>
      </div>
      <button type='button' @click="addNewLetterForCraftingRecipe" class='add-new-letter-button'>
        Add letter
        <svg viewBox="0 0 24 24" class='add-new-letter-button-icon'>
          <line x1="12" x2="12" y1="19" y2="5" />
          <line x1="5" x2="19" y1="12" y2="12" />
        </svg>
      </button>
    </div>
    <div class='crafting-table-grid-zone'>
      <div class='clear-letters-button' @click="clearLetters">Clear letters</div>

      <div class='crafting-table-grid-div-container'>
        <div class='crafting-table-grid-div'>
          <div class='crafting-table-grid-item' v-for="(item, index) in gridItems" :key="index"
            @dragover.prevent="handleDragOver" @drop.prevent="handleDrop($event, Math.floor(index / 3), index % 3)">
            {{ item }}
          </div>
        </div>
      </div>


      <DefLine labelText="output:">
        <InputWithSuggestions :value="initialOutput" @updateValue="this.outputValue = $event" suggestion-type="item" />
      </DefLine>
      <DefLine labelText="output count:" style="margin-top: 0;">
        <DefInputNum :value="initialOutputCount" @updateValue="this.outputCountValue = $event" />
      </DefLine>
    </div>
    <DefSave :submitText="submitText" />
  </form>
  <InvalidInputDialog ref="errDialog" :errorText="`${errDialogText}`"></InvalidInputDialog>
</template>
  
<script>
import DefSave from './../../../../default/DefSave.vue';
import DefCheckBox from './../../../../default/DefCheckBox.vue';
import DefLine from './../../../../default/DefLine.vue';
import DefInputNum from './../../../../default/DefInputNum.vue';
import InputWithSuggestions from './../../../../default/InputWithSuggestions.vue';
import InvalidInputDialog from './../../../../modalDialogs/InvalidInputDialog.vue';
import { CurrentProjectAddNewRecipe, CurrentProjectChangeAction } from "../../../../../../wailsjs/go/projectrelated/ProjectManager";


export default {
  props: {
    initialOutput: {
      type: String,
      default: ''
    },
    initialOutputCount: {
      type: Number,
      default: 1
    },
    isShapeless: {
      type: Boolean,
      default: false
    },
    letterItemDictionary: {
      type: Object,
      default: () => ({}),
    },
    gridLetters: {
      type: Array,
      default: () => [['', '', ''], ['', '', ''], ['', '', '']]
    },
    isNew: {
      type: Boolean,
      default: true
    },
    filePath: {
      type: String,
      default: ''
    },
    actionId: {
      type: Number,
      default: 0
    },
  },
  data() {
    return {
      outputValue: this.initialOutput,
      isShapelessValue: this.isShapeless,
      outputCountValue: this.initialOutputCount,
      errDialogText: '',
      letterItems: this.initLetterItems(),
      gridItems: this.initGridItems(),
    };
  },
  computed: {
    submitText() { return this.isNew ? 'Save to file' : 'Save changes'; }
  },
  methods: {
    handleSubmit() {
      this.errDialogText = "";
      const usedLettersInGrid = new Set(this.gridItems.filter(letter => letter.trim() !== ''));
      const definedLetters = new Set(this.letterItems.map(item => item.letter));
      const unusedLetters = [...definedLetters].filter(letter => !usedLettersInGrid.has(letter));
      const undefinedLetters = [...usedLettersInGrid].filter(letter => ![...definedLetters].some(definedLetter => definedLetter === letter));
      if (!this.outputValue) {
        this.errDialogText = "Output field is empty. ";
        this.$refs.errDialog.showDialog();
        return;
      }
      if (this.outputCountValue === null || this.outputCountValue === '') {
        this.errDialogText = "Output count field is empty. ";
        this.$refs.errDialog.showDialog();
        return;
      }
      if (this.outputCountValue > 128 || this.outputCountValue < 1) {
        this.errDialogText = "Output count cannot be more than 128 or less than 1";
        this.$refs.errDialog.showDialog();
        return;
      }
      if (unusedLetters.length > 0) {
        this.errDialogText = `The following letters are defined but not used in the grid: ${unusedLetters.join(', ')}`;
        this.$refs.errDialog.showDialog();
        return;
      }

      if (undefinedLetters.length > 0) {
        this.errDialogText = `The following letters are used in the grid but not defined: ${undefinedLetters.join(', ')}`;
        this.$refs.errDialog.showDialog();
        return;
      }
      const emptyValueLetters = this.letterItems.filter(item => item.value.trim() === '').map(item => item.letter);

      if (emptyValueLetters.length > 0) {
        this.errDialogText = `The following letters have empty values: ${emptyValueLetters.join(', ')}`;
        this.$refs.errDialog.showDialog();
        return;
      }
      const type = 'CraftingTableAdd';
      let formArgs = {
        letterItemDictionary: JSON.stringify(this.letterItems),
        lettersInputGrid: JSON.stringify(this.gridItems),
        isShapeless: this.isShapelessValue.toString(),
        output: this.outputValue,
        outputCount: this.outputCountValue.toString(),
      };
      console.log(formArgs);
      if (this.isNew) {
        let properties;
        console.log(type, formArgs);
        CurrentProjectAddNewRecipe(type, formArgs).then((historyItem) => {
          console.log(historyItem);
          properties = {
            initialInput: this.inputValue,
            initialOutput: this.outputValue,
            initialOutputCount: this.outputCountValue,
            filePath: historyItem.FilePath,
            actionId: historyItem.ActionID,
            isNew: false
          };
          this.newStoneCutterRecipeSaved(historyItem.ActionID, historyItem.ActionID, "new-recipe", properties);
        });
      }
      else {
        console.log("change");
        // CurrentProjectChangeAction(type, formArgs).then((historyItem) => {
        //   console.log(historyItem);
        // });
      }

    },
    initLetterItems() {
      return Object.entries(this.letterItemDictionary).map(([letter, value]) => ({
        letter,
        value,
      }));
    },
    addNewLetterForCraftingRecipe() {
      const currentChars = this.letterItems.map(item => item.letter);
      let nextChar = null;
      for (let i = 0; i < 9; i++) {
        const potentialChar = String.fromCharCode('a'.charCodeAt(0) + i).toUpperCase();
        if (!currentChars.includes(potentialChar)) {
          nextChar = potentialChar;
          break;
        }
      }

      if (!nextChar || this.letterItems.length >= 9) {
        this.errDialogText = 'You can`t add more than 9 letters';
        this.$refs.errDialog.showDialog();
        return;
      }

      this.letterItems.push({
        letter: nextChar,
        value: this.letterItemDictionary[nextChar] || '',
      });
    },
    removeLetter(index) {
      this.letterItems.splice(index, 1);
      this.errDialogText = '';
    },
    initGridItems() {
      return this.gridLetters.flat();
    },
    handleDragOver(event) {
      event.preventDefault();
    },
    handleDrop(event, i, j) {
      event.preventDefault();
      const letter = event.dataTransfer.getData('text/plain');
      this.gridItems[i * 3 + j] = letter;
    },

    handleDragStart(event, letter) {
      event.dataTransfer.setData('text/plain', letter);
    },
    clearLetters() {
      this.gridItems = this.gridItems.map(() => '');
    }
  },
  watch: {
    gridLetters: {
      handler(newVal) {
        this.gridItems = this.flattenGridLetters(newVal);
      },
      deep: true
    }
  },
  inject: ['newCraftingTableRecipeSaved'],
  components: {
    DefLine,
    InputWithSuggestions,
    DefSave,
    InvalidInputDialog,
    DefInputNum,
    DefCheckBox
  }
}
</script>

<style>
.item-for-letter-input .default-input {
  width: 91% !important;
}
</style>

<style scoped>
.form-container {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1.1fr;
  grid-gap: calc(4px + 1vw + 1vh);

}

.crafting-table-letters-zone {
  width: 100%;
  display: grid;
  grid-template-rows: calc(26px + 2vh + 0.5vw) 1fr calc(26px + 2vh + 0.5vw);
  position: relative;
  top: calc(24px + 4vh + 0.5vw);
  height: calc(90% - 4vh - 0.5vw - 36px);
  max-height: calc(98vh - 40px - 2vw);

}

.crafting-table-grid-zone {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.clear-letters-button {
  margin-top: 0;
  width: calc(4vw + 56px + 2vh);
  background-color: var(--back);
  border: 1px solid transparent;
  padding: calc(6px + 0.08vw + 0.17vh);
  border-radius: calc(5px + 0.11vw + 0.11vh);
  font-family: 'Figtree';
  font-size: calc(0.50vh + 0.50vw + 9px);
  font-weight: 300;
  color: var(--front);
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--back-3);
  transition: 0.08s;
}

.clear-letters-button:hover {
  transform: scale(1.015);
  background-color: var(--bright);
}

.crafting-table-grid-div-container {
  margin-top: calc(2vh - 2px);
  height: calc(55vh - 20px);
  max-width: calc(42vw - 2vh);
  aspect-ratio: 1/1;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: calc(2vh - 2px);
}

.crafting-table-grid-div {
  width: 100%;
  background-color: var(--front-2);
  aspect-ratio: 1/1;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(3, 1fr);
  gap: calc(4px + 0.13vw + 0.13vh);
  border-radius: calc(0.45vh + 0.45vw + 2px);
  border: calc(3px + 0.13vw + 0.13vh) solid var(--front-2);
}

.crafting-table-grid-div>div {
  align-self: center;
  justify-self: center;
  background-color: var(--back);
  width: 100%;
  height: 100%;
  border: 1px solid transparent;
  align-items: center;
  justify-content: center;
  display: flex;
  color: var(--front);
  font-family: 'Figtree';
  font-size: 180%;
}

.crafting-table-grid-div>div:nth-child(1) {
  border-top-left-radius: calc(0.45vh + 0.45vw + 1px);
}

.crafting-table-grid-div>div:nth-child(3) {
  border-top-right-radius: calc(0.45vh + 0.45vw + 1px);
}

.crafting-table-grid-div>div:nth-child(7) {
  border-bottom-left-radius: calc(0.45vh + 0.45vw + 1px);
}

.crafting-table-grid-div>div:nth-child(9) {
  border-bottom-right-radius: calc(0.45vh + 0.45vw + 1px);
}

.crafting-table-letters-container {
  overflow-y: auto;
  height: 100%;

}

.crafting-table-letters-container::-webkit-scrollbar {
  width: calc(0.32vw + 0.3vh + 2px);
}

.crafting-table-letters-container::-webkit-scrollbar-track {
  background: var(--back-3);
  border-radius: calc(0.1vw + 0.04vh + 2px);
}

.crafting-table-letters-container::-webkit-scrollbar-thumb {
  background: var(--front-2);
  border-radius: calc(0.1vw + 0.04vh + 2px);
}

.crafting-table-letters-container::-webkit-scrollbar-thumb:hover {
  background: var(--bright);
}

.crafting-table-letter-item {
  height: calc(0.6vw + 4vh + 20px);
  display: grid;
  grid-template-columns: 2fr 8fr 1fr;
  align-items: center;
  border: 1px solid transparent;
  border-radius: calc(2px + 0.08vw + 0.11vh);
}

.letter-delete-button-container {
  height: calc(12px + 36%);
  padding: 2px;
  max-height: calc(100%);
  aspect-ratio: 1/1;
  border: 0px solid transparent;
  border-radius: 30%;
  cursor: pointer;
}

.letter-delete-button-container:hover {
  background-color: var(--back-3);
}

.letter-delete-button path {
  stroke: var(--front);
}

.letter-label {
  font-family: 'Figtree';
  color: var(--front);
  align-self: center;
  justify-self: center;
  font-size: calc(0.57vh + 0.57vw + 11px);
}

.add-new-letter-button {
  justify-self: center;
  width: 92%;
  height: 100%;
  background-color: var(--bright-2);
  border: 1px solid transparent;
  border-radius: calc(3px + 0.09vw + 0.1vh);
  font-family: 'Figtree';
  font-size: calc(0.57vh + 0.57vw + 13px);
  font-weight: 300;
  color: var(--front);
  cursor: pointer;
  transition: 0.06s;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.add-new-letter-button:hover {
  transform: scale(1.024);
  background-color: var(--bright-3);
}

.add-new-letter-button .add-new-letter-button-icon {
  padding-top: calc(1px + 0.4%);
  height: calc(6px + 47%);
  aspect-ratio: 1/1;
  stroke: var(--front);
}

.add-new-letter-button .add-new-letter-button-icon line {
  stroke-width: 3;
}

.is-shapeless-checkbox {
  margin-left: calc(13px + 4%);
  margin-top: calc(0.2vw + 0.5vh);
}
</style>
