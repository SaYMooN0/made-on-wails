<template>
    <div class="page-container">
        <h1 class="new-theme-name">
            <div style="display: flex;justify-content: left; gap: 1vw;align-items: center;">
                New theme name:
                <input type="text" class="theme-name-input" v-model="themeName" />
            </div>
            <div class="save-theme-button" @click="saveThemes">Save</div>
        </h1>
        <div class="color-inputs-blocks-container">
            <ColorInputsBlock v-if="Object.keys(newTheme).length > 0" ref="backColorsBlock" title="Back Colors" type="Back"
                :colorList="[
                    { label: 'Main', color: newTheme.BackMainColor },
                    { label: 'Second', color: newTheme.BackSecondColor },
                    { label: 'Third', color: newTheme.BackThirdColor }
                ]" />
            <ColorInputsBlock v-if="Object.keys(newTheme).length > 0" ref="frontColorsBlock" title="Front Colors"
                type="Front" :colorList="[
                    { label: 'Main', color: newTheme.FrontMainColor },
                    { label: 'Second', color: newTheme.FrontSecondColor },
                    { label: 'Third', color: newTheme.FrontThirdColor }
                ]" />
            <ColorInputsBlock v-if="Object.keys(newTheme).length > 0" ref="brightColorsBlock" title="Bright Colors"
                type="Bright" :colorList="[
                    { label: 'Main', color: newTheme.BrightMainColor },
                    { label: 'Second', color: newTheme.BrightSecondColor },
                    { label: 'Third', color: newTheme.BrightThirdColor }
                ]" />
            <ColorInputsBlock v-if="Object.keys(newTheme).length > 0" ref="warningColorsBlock" title="Warning Colors"
                type="Warning" :colorList="[
                    { label: 'Main', color: newTheme.WarningMainColor },
                    { label: 'Bright', color: newTheme.WarningBrightColor }
                ]" />
        </div>
        <GoBackButton @goBack="backToMain" />
    </div>
    <ErrorDialog ref="errDialog" :errorText="`${errDialogText}`" />
</template>
<script>
import ColorInputsBlock from '../newThemePageComponents/ColorInputsBlock.vue'
import GoBackButton from '../GoBackButton.vue';
import { GetHexDefaultDark, GetAllThemeNames, AddNewTheme, ThemeFromHexTheme } from '../../../wailsjs/go/themerelated/ThemeCollection';
import ErrorDialog from '../modalDialogs/ErrorDialog.vue';
export default {
    data() {
        return {
            newTheme: {},
            allThemeNames: [],
            themeName: 'new theme',
            errDialogText: ''
        }

    },
    computed: {
        themeValues() {
            return Object.fromEntries(
                Object.entries(this.newTheme).filter(([key]) => key.endsWith("Color"))
            );
        },
    },
    components:
    {
        ColorInputsBlock,
        GoBackButton,
        ErrorDialog
    },
    methods: {
        backToMain() {
            this.$emit('goBack');
        },
        saveThemes() {
            if (this.allThemeNames.includes(this.themeName)) {
                this.errDialogText = "You already have a theme with that name. Please come up with another name for this theme";
                this.$refs.errDialog.showDialog();
                return;
            }
            const backColors = this.$refs.backColorsBlock.getColors();
            const frontColors = this.$refs.frontColorsBlock.getColors();
            const brightColors = this.$refs.brightColorsBlock.getColors();
            const warningColors = this.$refs.warningColorsBlock.getColors();
            const newThemeName = this.themeName;
            const newThemeValues = {
                ...backColors,
                ...frontColors,
                ...brightColors,
                ...warningColors,
                Name:newThemeName
            };

            ThemeFromHexTheme(newThemeValues).then((themeRGB) => {
                console.log(themeRGB);
                // AddNewTheme(themeRGB);
            });
        }
    },
    created() {
        GetHexDefaultDark().then((theme) => {
            this.newTheme = theme;
        });
        GetAllThemeNames().then((names) => {
            this.allThemeNames = names;
        });
    }
};
</script>
<style>
.page-container {
    display: grid;
    grid-template-rows: auto 1fr;
    height: 100vh;
    padding-left: 4vw;
    padding-right: 4vw;
}

.new-theme-name {
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(1vh + 1.6vw + 11px);
    margin-top: calc(4vh + 2vw + 10px);
    display: flex;
    justify-content: space-between;
    align-items: center;

}

.theme-name-input {
    background-color: var(--back-2);
    border: 1px solid transparent;
    border-radius: calc(2px + 0.1vw + 0.1vh);
    font-family: 'Figtree';
    font-size: calc(0.8vh + 1vw + 9px);
    font-weight: 300;
    color: var(--front);
    padding-left: calc(1px + 0.2vw + 0.2vh);
    padding-right: calc(1px + 0.2vw + 0.2vh);
    width: 44%;
}

.theme-name-input:focus {
    outline: none;
    background-color: var(--back-3);
}

.save-theme-button {
    height: calc(10px + 50%);
    background-color: var(--bright-2);
    border: 1px solid transparent;
    border-radius: calc(3px + 0.09vw + 0.1vh);
    font-family: 'Figtree';
    font-size: calc(0.5vh + 0.6vw + 10px);
    color: var(--front);
    padding-left: calc(5px + 0.3vw + 0.3vh);
    padding-right: calc(5px + 0.3vw + 0.3vh);
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    transition: 0.08s;
}

.save-theme-button:hover {
    transform: scale(1.02);
    background-color: var(--bright-3);
}

.color-inputs-blocks-container {
    height: 90%;
    width: 100%;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: calc(0.4vh + 1vw);
    align-items: center;
}
</style>