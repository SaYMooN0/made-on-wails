<template>
  <form @submit.prevent="handleSubmit" class="form-container">
    <DefLine labelText="input:">
      <InputWithSuggestions :value="initialInput" @updateValue="this.inputValue = $event" suggestion-type="item" />
    </DefLine>
    <DefLine labelText="output:">
      <InputWithSuggestions :value="initialOutput" @updateValue="this.outputValue = $event" suggestion-type="item" />
    </DefLine>
    <DefLine labelText="output count:">
      <DefInputNum :value="initialOutputCount" @updateValue="this.outputCountValue = $event" />
    </DefLine>
    <DefSave :submitText="submitText" />
  </form>
  <InvalidInputDialog ref="errDialog" :errorText="`${errDialogText}`"/>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import InputWithSuggestions from './../../../../default/InputWithSuggestions.vue';
import DefInputNum from './../../../../default/DefInputNum.vue';
import DefSave from './../../../../default/DefSave.vue';
import InvalidInputDialog from './../../../../modalDialogs/InvalidInputDialog.vue';
import { CurrentProjectAddNewRecipe, CurrentProjectChangeAction } from "../../../../../../wailsjs/go/projectrelated/ProjectManager";
export default {
  props: {
    initialInput: {
      type: String,
      default: ''
    },
    initialOutput: {
      type: String,
      default: ''
    },
    initialOutputCount: {
      type: Number,
      default: 1
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
      inputValue: this.initialInput,
      outputValue: this.initialOutput,
      outputCountValue: this.initialOutputCount,
      errDialogText: '',
    };
  },
  computed: {
    submitText() {
      return this.isNew ? 'Save to file' : 'Save changes';
    }
  },
  methods: {
    handleSubmit() {
      this.errDialogText = "";

      if (!this.inputValue) {
        this.errDialogText = "Input field is empty. ";
        this.$refs.errDialog.showDialog();
        return;
      }
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
      const type = 'StoneCutterAdd';
      let formArgs = {
        input: this.inputValue,
        output: this.outputValue,
        outputCount: this.outputCountValue.toString()
      };

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

  },
  inject: ['newStoneCutterRecipeSaved'],
  components: {
    DefLine,
    InputWithSuggestions,
    DefInputNum,
    DefSave,
    InvalidInputDialog
  }
}
</script>

<style scoped>
.form-container {
  padding-top: calc(2vw + 10px + 20vh);
  padding-left: calc(4vw + 20px + 2vh);
  width: calc(100% - 4vw - 20px - 2vh);
  display: block;
}
</style>
