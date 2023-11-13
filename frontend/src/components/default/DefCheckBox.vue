<template>
  <label class="default-checkbox-label">
    {{ labelText }}
    <input type="checkbox" :checked="modelValue" @change="$emit('update:modelValue', $event.target.checked)">
    <span class="checkmark"></span>
  </label>
</template>
  
<script>
export default {
  props: {
    modelValue: {
      type: Boolean,
      default: false
    },
    labelText: {
      type: String,
      default: '' // Значение по умолчанию - пустая строка
    }
  },
  emits: ['update:modelValue']
};
</script>
  
<style>
.default-checkbox-label input {
  display: none;
}

.default-checkbox-label {
  color: var(--front);
  font-family: 'Figtree';
  font-weight: 600;
  font-size: calc(0.7vw + 0.7vh + 9px);
  white-space: nowrap;
}

.default-checkbox-label .checkmark {
  position: relative;
  top: 0;
  left: calc(0.1vw + 2px);
  aspect-ratio: 1/1;
  height: calc(11px + 0.3vw + 0.28vh);
  background-color: var(--back-2);
  margin-right: 10px;
  display: inline-block;
  vertical-align: middle;
  border: calc(2px + 0.05vw + 0.05vh) solid transparent;
  border-radius: 10%;
  border: 2px solid var(--front-2);
}

.default-checkbox-label:hover input~.checkmark {
  background-color: var(--back-3);
  border-color: var(--front-3);
}

.default-checkbox-label input:checked~.checkmark {
  background-color: var(--front-3);
  border-color: var(--front-3);
}

.default-checkbox-label .checkmark:after {
  content: "";
  position: absolute;
  display: none;

}

.default-checkbox-label input:checked~.checkmark:after {
  display: block;
}

.default-checkbox-label .checkmark:after {
  width: 28%;
  height: 70%;
  border: solid white;
  border-width: 0 3px 3px 0;
  transform: rotate(45deg);
  left: 54%;
  top: calc(50% - 1px);
  transform: translate(-50%, -50%) rotate(45deg);
}
</style>