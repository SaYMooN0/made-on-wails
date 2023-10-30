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
  <InvalidInputDialog ref="errDialog" :errorText="`${errDialogText}`"></InvalidInputDialog>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import InputWithSuggestions from './../../../../default/InputWithSuggestions.vue';
import DefInputNum from './../../../../default/DefInputNum.vue';
import DefSave from './../../../../default/DefSave.vue';
import InvalidInputDialog from './../../../../modalDialogs/InvalidInputDialog.vue';
export default {
  props: {
    initialInput: {
      type: String,
      default: ''
    },
    initialOuput: {
      type: String,
      default: ''
    },
    initialOutputCount: {
      type: Number,
      default: 1
    },
    isNew: {
      type: Boolean,
      default: false
    }
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
      return this.isNew ? 'Save changes' : 'Save to file';
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
      const type = 'StonecutterAdd'
      let formArgs = {
        input: this.inputValue,
        output: this.outputValue,
        outputCount: this.outputCountValue
      };
      console.log(formArgs);
    },

  },
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
  padding-top: calc(2vw + 10px + 10vh);
  padding-left: calc(4vw + 20px + 2vh);
  width: 100%;
  display: block;
}
</style>
