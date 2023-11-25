<template>
    <form @submit.prevent="handleSubmit" class="form-container">
        <DefLine labelText="input:">
            <InputWithSuggestions :value="initialInput" @updateValue="this.inputValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine labelText="output:">
            <InputWithSuggestions :value="initialOutput" @updateValue="this.outputValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine labelText="additional:">
            <div class="radio-btns-container">
                <DefRadio name="additional" value="FurnaceOnly" spanText="None" v-model="furnaceTypeValue" />
                <DefRadio name="additional" value="FurnaceAndSmoker" spanText="Smoker" v-model="furnaceTypeValue" />
                <DefRadio name="additional" value="FurnaceAndBlast" spanText="Blast Furnace" v-model="furnaceTypeValue" />
            </div>
        </DefLine>
        <DefSave :submitText="submitText" />
    </form>
    <ErrorDialog ref="errDialog" :errorText="`${errDialogText}`"/>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import InputWithSuggestions from './../../../../default/InputWithSuggestions.vue';
import DefRadio from './../../../../default/DefRadio.vue';
import DefSave from './../../../../default/DefSave.vue';
import ErrorDialog from './../../../../modalDialogs/ErrorDialog.vue';
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
        initialFurnaceType: {
            type: String,
            default: "FurnaceOnly"
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
            furnaceTypeValue: this.initialFurnaceType,
            errDialogText: '',
        };
    },
    computed: {
        isFurnaceOnlyChecked() { return this.initialFurnaceType === "FurnaceOnly"; },
        isFurnaceAndSmokerChecked() { return this.initialFurnaceType === "FurnaceAndSmoker"; },
        isFurnaceAndBlastChecked() { return this.initialFurnaceType === "FurnaceAndBlast"; },
        submitText() { return this.isNew ? 'Save to file' : 'Save changes'; }
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
            console.log(this.furnaceTypeValue);
            if (this.furnaceTypeValue===undefined) {
                this.errDialogText = "Choose special furnace type. ";
                this.$refs.errDialog.showDialog();
                return;
            }
            const type = this.furnaceTypeValue+'Add';
            let formArgs = {
                input: this.inputValue,
                output: this.outputValue,
                specialType: this.type
            };
            if (this.isNew) {
                let properties;
                CurrentProjectAddNewRecipe(type, formArgs).then((historyItem) => {
                    console.log(historyItem);
                    properties = {
                        initialInput: this.inputValue,
                        initialOutput: this.outputValue,
                        initialFurnaceType: this.furnaceTypeValue,
                        filePath: historyItem.FilePath,
                        actionId: historyItem.ActionID,
                        isNew: false
                    };
                    this.newFurnaceRecipeSaved(historyItem.ActionID, historyItem.ActionID, "new-recipe", properties);
                });
            }
            else {
                console.log("change");
                // CurrentProjectChangeAction(type, formArgs).then((historyItem) => {
                //   console.log(historyItem);
                // });
            }
        }
    },
    inject: ['newFurnaceRecipeSaved'],
    components: {
        DefLine,
        InputWithSuggestions,
        DefRadio,
        DefSave,
        ErrorDialog
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

.radio-btns-container {
    width: 60%;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    justify-content: right;
    gap: 10px;
}

@media (max-width: 850px) {
    .radio-btns-container {
        grid-template-columns: 1fr;
    }
}
</style>
  