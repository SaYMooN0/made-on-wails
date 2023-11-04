<template>
  <div>
    <input type="text" class="default-input" @input="handleInput" ref="inputRef"
      @focus="showSuggestions = true; activeIndex = 0" v-model="inputValue" />
    <div v-if="showSuggestions && filteredSuggestions.length" class="suggestions-list">
      <div v-for="(suggestion, index) in filteredSuggestions" :key="suggestion" @click.stop="setActiveSuggestion(index)"
        @dblclick="selectSuggestion(suggestion)" :ref="`suggestion-${index}`"
        :class="{ 'active': activeIndex === index }">
        {{ suggestion }}
      </div>
    </div>
  </div>
</template>

<script>
import { CurrentProjectGetItemsTypeSuggestion } from "../../../wailsjs/go/projectrelated/ProjectManager";
export default {
  props: {
    suggestionType: {
      type: String,
      default: 'none'
    },
    value: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      inputValue: this.value,
      showSuggestions: false,
      activeIndex: 0,
      suggestionItems: []
    };
  },
  watch: {
    inputValue: {
      handler: function (newValue) {
        this.suggestionItems=[];
        if (this.suggestionType == "item") {
          CurrentProjectGetItemsTypeSuggestion(newValue.toLowerCase()).then((resievedSuggestions) => {
            this.suggestionItems = resievedSuggestions;
          });
        }
      },
      immediate: true
    }
  },
  computed: {
    filteredSuggestions() {
      if (this.suggestionType == "type") {
        return ["typqdf", "dsdsa", "dsda", "iopjqwdopdsop", "cczxczxc"];
      }
      return this.suggestionItems;
    }
  },
  methods: {
    handleInput(event) {
      this.showSuggestions = true;
      this.inputValue = event.target.value;
      this.$emit('updateValue', this.inputValue);
    },
    ensureActiveItemVisible() {
      const activeItem = this.$refs[`suggestion-${this.activeIndex}`][0];
      if (activeItem) {
        const list = activeItem.parentElement;
        const itemTop = activeItem.offsetTop;
        const itemBottom = itemTop + activeItem.offsetHeight;

        if (itemTop < list.scrollTop) {
          list.scrollTop = itemTop;
        } else if (itemBottom > list.scrollTop + list.clientHeight) {
          list.scrollTop = itemBottom - list.clientHeight;
        }
      }
    },
    setActiveSuggestion(index) {
      this.activeIndex = index;
    },
    selectSuggestion(suggestion) {
      this.showSuggestions = false;
      this.inputValue = suggestion;
      this.$emit('updateValue', this.inputValue);
    },
    handleOutsideClick(event) {
      const inputContains = this.$refs.inputRef && this.$refs.inputRef.contains(event.target);
      const suggestionsContains = this.$refs.suggestionsRef && this.$refs.suggestionsRef.contains(event.target);

      if (!inputContains && !suggestionsContains) {
        this.showSuggestions = false;
      }
    },
    handleKeydown(event) {
      if (this.showSuggestions) {
        if (this.showSuggestions) {
          switch (event.key) {
            case 'ArrowDown':
              if (this.activeIndex < this.filteredSuggestions.length - 1) {
                this.activeIndex++;
              } else {
                this.activeIndex = 0;
              }
              break;
            case 'ArrowUp':
              if (this.activeIndex > 0) {
                this.activeIndex--;
              } else {
                this.activeIndex = this.filteredSuggestions.length - 1;
              }
              break;
            case 'Enter':
            case 'Tab':
              if (this.filteredSuggestions.length > 0) {
                event.preventDefault();
                this.selectSuggestion(this.filteredSuggestions[this.activeIndex]);
              }
              break;
            case 'Escape':
              this.showSuggestions = false;
              break;
          }
          this.$nextTick(() => {
            this.ensureActiveItemVisible();
          });
        }
      }
    }
  },
  mounted() {
    document.addEventListener("click", this.handleOutsideClick);
    document.addEventListener("keydown", this.handleKeydown);  // Добавьте эту строку
  },
  beforeDestroy() {
    document.removeEventListener("click", this.handleOutsideClick);
    document.removeEventListener("keydown", this.handleKeydown);  // Добавьте эту строку
  },
}
</script>

<style scoped>
.suggestions-list {
  position: absolute;
  max-height: calc(80px + 8vh);
  overflow-y: auto;
  width: 200px;
  z-index: 1000;
  background-color: var(--back);
  border: 1px solid var(--back-3);
  color: var(--front);
  font-family: Bahnschrift;
  font-weight: 600;
  font-size: calc(0.36vh + 0.36vw + 9px);
  margin-left: 16px;
}

.suggestions-list .active {
  background-color: var(--back-3);
  color: var(--front-3);
}

.suggestions-list div:hover {
  background-color: var(--back-2);
}

.suggestions-list::-webkit-scrollbar {
  width: calc(0.15vw + 5px);
  max-width: calc(6px);
}

.suggestions-list::-webkit-scrollbar-track {
  background-color: var(--back-2);
}

.suggestions-list::-webkit-scrollbar-thumb {
  background-color: var(--front-2);
  height: calc(2vh + 5px);
  border: 1px solid transparent;
  border-radius: calc(2px + 0.12vw + 0.12vh);
}

.suggestions-list::-webkit-scrollbar-thumb:hover {
  background-color: var(--bright);
}

.default-input {
  height: calc(1.6vh + 8px + 0.25vw);
  width: calc(20vw + 45px);
  background-color: var(--back-2);
  border: 1px solid transparent;
  border-radius: calc(2px + 0.05vw + 0.1vh);
  font-family: 'Figtree';
  font-size: calc(0.88vh + 0.6vw + 7px);
  font-weight: 300;
  color: var(--front);
  padding-left: calc(1px + 0.2vw + 0.2vh);
  padding-right: calc(1px + 0.2vw + 0.2vh);
  position: relative;
}

.default-input:focus {
  outline: none;
  background-color: var(--back-3);
}
</style>
