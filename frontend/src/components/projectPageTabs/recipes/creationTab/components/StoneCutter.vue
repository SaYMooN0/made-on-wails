<template>
  <form @submit.prevent="handleSubmit" class="form-container">
    <DefLine labelText="input:">
      <DefInput ref="inputRef" :value="inputRef" suggestion="item" />
    </DefLine>
    <DefLine labelText="output:">
        <DefInput ref="outputRef" :value="outputRef" suggestion="item" />
    </DefLine>
    <DefLine labelText="output count:">
        <DefInputNum ref="outputCountRef" :value="outputCountRef"/>
    </DefLine>
    <DefSave @click="logValues" :submitText="submitTextValue"/>
  </form>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import DefInput from './../../../../default/DefInput.vue';
import DefInputNum from './../../../../default/DefInputNum.vue';
import DefSave from './../../../../default/DefSave.vue';

export default {
  data() {
    return {
      inputRef: this.inputValue,
      outputRef: this.outputValue,
      outputCountRef: this.outputCountValue
    }
  },
  props: {
    formArguments: Object,
    actionId: [String, Number],
    path: String,
    inputValue: {
      type: String,
      default: ''
    },
    outputValue: {
      type: String,
      default: ''
    },
    outputCountValue: {
      type: Number,
      default: 1
    },
    submitTextValue: {
      type: String,
      default: 'Save'
    }
  },
  methods: {
    handleSubmit() {
    },
    logValues() {
      // console.log(this.input);
      console.log(this.$refs.inputRef.$el.value);
      console.log(this.$refs.outputRef.$el.value);
      console.log(this.$refs.outputCountRef.$el.value);
    }
  },
  components: {
    DefLine,
    DefInput,
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
