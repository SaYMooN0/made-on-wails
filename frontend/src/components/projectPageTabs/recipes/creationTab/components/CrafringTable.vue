<template>
  <form @submit.prevent="handleSubmit" class="form-container">
    <div class='crafting-table-letters-zone'>
    <DefCheckBox :name="isShapeless" class="is-shapeless-checkbox"/>
       <!-- <div class='crafting-table-letters-container'>${letterDivs}</div>
      <button type='button' onclick="addNewLetterForCraftingRecipe(event)" class='add-new-letter-button'>
        Add letter
        <svg viewBox="0 0 24 24" class='add-new-letter-button-icon'>
          <line x1="12" x2="12" y1="19" y2="5" />
          <line x1="5" x2="19" y1="12" y2="12" />
        </svg>
      </button> -->
    </div>
    <div class='crafting-table-grid-zone'>
      <div class='clear-letters-button' onclick='clearLetters()'>Clear letters</div>
      <div class='crafting-table-grid-div-container'>
        <div class='crafting-table-grid-div'>
          ${gridItems}
        </div>
      </div>
      <div class="crafting-table-output-div">
        <DefLine labelText="output:">
          <InputWithSuggestions :value="initialOutput" @updateValue="this.outputValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine labelText="output count:">
          <DefInputNum :value="initialOutputCount" @updateValue="this.outputCountValue = $event" />
        </DefLine>
      </div>
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
      isShapelessValue: this.isShapelessType,
      outputCountValue: this.initialOutputCount,
      errDialogText: '',
    };
  },
  computed: {
    submitText() { return this.isNew ? 'Save to file' : 'Save changes'; }
  },
  methods: {
    handleSubmit(event) {

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
  grid-template-rows: auto 1fr calc(14px + 4vh + 0.8vw);
  grid-gap: calc(0.8vh + 2px);
  background-color: darkblue;
}
.is-shapeless-checkbox {
    margin-left: calc(13px + 4%);
    margin-top:calc(0.2vw + 0.5vh);
}
.crafting-table-grid-zone {
  width: 100%;
  height: 100%;
  display: flex; 
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: cadetblue;
}

.clear-letters-button {
  margin-top: -2vw;
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
  margin-top: calc(10px + 2%);
  height: calc(63vh - 3vw - 40px);
  max-width: calc(47vw - 2vh);
  aspect-ratio: 1/1;
  display: flex;
  justify-content: center;
  align-items: center;
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

.item-for-letter-input {
  width: 94%;
  height: calc(6px + 44%);
  border: 1px solid transparent;
  border-radius: calc(1px + 0.04vw + 0.1vh);
  color: var(--front-3);
  font-family: 'Bahnschrift';
  background-color: transparent;
  border: none;
  outline: none;
  font-size: calc(1vh + 0.12vw + 8px);
}

.item-for-letter-input:focus {
  outline: none;
  background-color: var(--back-3);
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
  width: calc(98% - 10px);
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

.crafting-table-output-div {
  margin-top: calc(2% + 10px);
  display: grid;
  width: 100%;
  grid-template-rows: 1fr 1fr;
  gap: calc(2% + 10px);
  height: calc(30px + 10%);
}

.is-shapeless-checkbox {
  margin-left: calc(13px + 4%);
  margin-top: calc(0.2vw + 0.5vh);
}
</style>
  