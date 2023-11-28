<template>
  <h1>Madeinfo</h1>
  <h1>Madeinfo</h1>
  <h1>Madeinfo</h1>
  <h1>Madeinfo</h1>
  <h1>Madeinfo</h1>
  <h1>Madeinfo</h1>
  <h1>Madeinfo</h1>
  <GoBackButton @goBack="backToMain"/>
</template>

<script>
import GoBackButton from '../GoBackButton.vue';
export default {
  components:
  {
    GoBackButton
  },
  methods: {
    backToMain() {
      this.$emit('goBack');
    }
  }

}
</script>