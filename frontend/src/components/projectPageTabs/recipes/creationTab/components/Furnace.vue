<template>
    <form @submit.prevent="handleSubmit" class="furnace-form">
        <p class="input-line">
            <label class="default-input-label"> input: </label>
            <input class="default-input" type="text" v-model="inputValue">
        </p>
        <p class="input-line">
            <label class="default-input-label"> output: </label>
            <input class="default-input" type="text" v-model="outputValue">
        </p>
        <div class="furnace-radios-container">
            <label class="furnace-additional-type-label">Additional</label>
            <label v-for="type in radioTypes" :key="type.value" class="default-radio-container">
                <input class="default-radio" type="radio" name="furnace-type-choice" :value="type.value"
                    v-model="specialTypeValue">
                <span class="default-radio-label">{{ type.label }}</span>
            </label>
        </div>
        <label class="default-error-label"></label>
        <p class="input-line">
            <input class="default-submit" type="submit" :value="submitButtonText">
        </p>
    </form>
</template>
  
<script>
export default {
    props: ['formArguments', 'actionId', 'path'],
    data() {
        return {
            radioTypes: [
                { value: "FurnaceOnly", label: "None" },
                { value: "FurnaceAndSmoker", label: "Smoker" },
                { value: "FurnaceAndBlast", label: "Blast Furnace" }
            ]
        };
    },
    computed: {
        inputValue() {
            return this.formArguments && this.formArguments.input ? this.formArguments.input : '';
        },
        outputValue() {
            return this.formArguments && this.formArguments.output ? this.formArguments.output : '';
        },
        specialTypeValue: {
            get() {
                return this.formArguments && this.formArguments.specialType ? this.formArguments.specialType : '';
            },
            set(value) {
                this.$emit('update:specialTypeValue', value);
            }
        },
        submitButtonText() {
            return this.formArguments ? "Save changes" : "Save to file";
        }
    },
    methods: {
        handleSubmit(event) {
        }
    }
}
</script>
  
<style scoped></style>
  