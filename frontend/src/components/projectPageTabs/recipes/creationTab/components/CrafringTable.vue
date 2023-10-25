<template>
    <form @submit.prevent="handleSubmit" class="crafting-table-form">
      <div class="crafting-grid">
        <div v-for="row in 3" :key="'row' + row" class="crafting-row">
          <input 
            v-for="col in 3" 
            :key="'cell' + col" 
            class="crafting-cell" 
            v-model="craftingRecipe[`${row}-${col}`]"
            placeholder="Item">
        </div>
      </div>
      <div class="recipe-output">
        <label for="outputItem">Output Item:</label>
        <input id="outputItem" type="text" v-model="outputItem" placeholder="Crafted Item">
      </div>
      <div>
        <input class="crafting-submit" type="submit" :value="submitButtonText">
      </div>
    </form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        craftingRecipe: {},
        outputItem: ''    
      };
    },
    computed: {
      submitButtonText() {
        return this.craftingRecipe && Object.keys(this.craftingRecipe).length > 0 ? "Update Recipe" : "Save Recipe";
      }
    },
    methods: {
      handleSubmit(event) {
        // Your form submission logic here...
        // Maybe emit an event with the crafted recipe to the parent component
        this.$emit('crafting-recipe-submitted', {
          recipe: this.craftingRecipe,
          output: this.outputItem
        });
      }
    }
  }
  </script>
  
  <style scoped>
  /* Any styling specific to this component */
  .crafting-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }
  
  .crafting-cell {
    width: 50px;
    height: 50px;
    padding: 5px;
  }
  
  .recipe-output {
    margin-top: 20px;
  }
  </style>
  