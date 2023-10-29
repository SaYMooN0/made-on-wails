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
    <DefSave :submitText="submitTextValue" />
  </form>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import InputWithSuggestions from './../../../../default/InputWithSuggestions.vue';
import DefInputNum from './../../../../default/DefInputNum.vue';
import DefSave from './../../../../default/DefSave.vue';

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
    submitTextValue: {
      type: String,
      default: 'Save'
    }
  },
  data() {
    return {
      inputValue: this.initialInput,
      outputValue: this.initialOutput,
      outputCountValue: this.initialOutputCount,
      submitText: this.submitTextValue
    };
  },
  methods: {
    handleSubmit() {
      console.log(this.inputValue);
      console.log(this.outputValue);
      console.log(this.outputCountValue);
    },

  },
  components: {
    DefLine,
    InputWithSuggestions,
    DefInputNum,
    DefSave
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
