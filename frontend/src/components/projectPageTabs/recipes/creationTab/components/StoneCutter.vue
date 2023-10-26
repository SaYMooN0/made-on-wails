<template>
  <form @submit.prevent="handleSubmit">
    <DefLine labelText="input:">
      <DefInput />
    </DefLine>
    <DefLine labelText="output:">
      <DefInput />
      <input class="default-input-num" type="number">
    </DefLine>
  </form>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import DefInput from './../../../../default/DefInput.vue';

export default {
  props: ['formArguments', 'actionId', 'path'],
  computed: {
    inputValue() {
      return this.formArguments && this.formArguments.input ? this.formArguments.input : '';
    },
  },
  methods: {
    handleSubmit() {
    }
  },
  components: {
    DefLine,
    DefInput
  }
}
</script>
  