<template>
    <div class="visual-settings-container">
        <div class="chosen-theme-zone">
            <label class="chosen-theme-name">Theme: {{ chosenTheme.Name }}</label>
            <div class="color-labels-container">
                <ColorLabel v-for="(value, key) in chosenThemeColorFields" :key="key" :label="key" :color="value"
                    ref="colorLabels" />
            </div>
        </div>
        <div class="avaliable-themes-zone">
            <p class="avaliable-themes-label">Available themes:</p>
            <div class="themes-labels-container scrollbar">
                <ThemeLabel v-for="theme in avaliableThemes" :theme="theme" @click="changeTheme(theme)" />
            </div>
            <button class="new-theme-button" @click="createNewTheme">
                Create new theme
                <svg viewBox="0 0 24 24" fill="none" class="brush-icon">
                    <path
                        d="M9.5 19.5V18H4.5C3.95 18 3.45 17.78 3.09 17.41C2.72 17.05 2.5 16.55 2.5 16C2.5 14.97 3.3 14.11 4.31 14.01C4.37 14 4.43 14 4.5 14H19.5C19.57 14 19.63 14 19.69 14.01C20.17 14.05 20.59 14.26 20.91 14.59C21.32 14.99 21.54 15.56 21.49 16.18C21.4 17.23 20.45 18 19.39 18H14.5V19.5C14.5 20.88 13.38 22 12 22C10.62 22 9.5 20.88 9.5 19.5Z"
                        stroke="#292D32" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
                    <path
                        d="M20.17 5.3L19.69 14.01C19.63 14 19.57 14 19.5 14H4.50001C4.43001 14 4.37001 14 4.31001 14.01L3.83001 5.3C3.65001 3.53 5.04001 2 6.81001 2H17.19C18.96 2 20.35 3.53 20.17 5.3Z"
                        stroke="#292D32" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
                    <path d="M7.98999 2V7" stroke="#292D32" stroke-width="1.5" stroke-linecap="round"
                        stroke-linejoin="round" />
                    <path d="M12 2V4" stroke="#292D32" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
                </svg>
            </button>
        </div>
    </div>
    <ErrorDialog ref="errDialog" :errorText="`${errDialogText}`" />
</template>
<script>


import { CurrentThemeColors, GetAllThemesValues, UpdateTheme, ThemeFromHexTheme, SetCurrentTheme } from "./../../../../wailsjs/go/themerelated/ThemeCollection";
import ColorLabel from "../tabsComponents/ColorLabel.vue";
import ThemeLabel from "../tabsComponents/ThemeLabel.vue";
import ErrorDialog from "../../modalDialogs/ErrorDialog.vue"

export default {
    data() {
        return {
            avaliableThemes: [],
            chosenTheme: {},
            errDialogText: '',
        };
    },
    methods: {
        changeTheme(theme) {
            this.cancelAllEditing();
            SetCurrentTheme(theme.Name).then((returnedString) => {
                if (returnedString) {
                    this.showError(returnedString);
                }
                else {
                    this.chosenTheme = theme;
                    this.applyCurrentThemeColors();
                }
            });
        }
        ,
        applyCurrentThemeColors() {
            CurrentThemeColors().then((theme) => {
                document.documentElement.style.setProperty('--back', theme.BackMainColor);
                document.documentElement.style.setProperty('--back-2', theme.BackSecondColor);
                document.documentElement.style.setProperty('--back-3', theme.BackThirdColor);
                document.documentElement.style.setProperty('--front', theme.FrontMainColor);
                document.documentElement.style.setProperty('--front-2', theme.FrontSecondColor);
                document.documentElement.style.setProperty('--front-3', theme.FrontThirdColor);
                document.documentElement.style.setProperty('--bright', theme.BrightMainColor);
                document.documentElement.style.setProperty('--bright-2', theme.BrightSecondColor);
                document.documentElement.style.setProperty('--bright-3', theme.BrightThirdColor);
                document.documentElement.style.setProperty('--warning-main', theme.WarningMainColor);
                document.documentElement.style.setProperty('--warning-bright', theme.WarningBrightColor);
            });
        },
        cancelAllEditing() {
            if (this.$refs.colorLabels) {
                this.$refs.colorLabels.forEach(colorLabelComponent => {
                    colorLabelComponent.cancelEditing();
                });
            }
        },
        updateTheme(themeField, colorValue) {
            this.chosenTheme[themeField] = colorValue;
            ThemeFromHexTheme(this.chosenTheme).then((validThemeValues) => {
                UpdateTheme(validThemeValues.Name, validThemeValues).then((returnedString) => {
                    if (returnedString) {
                        this.showError(returnedString);
                    }
                    else {
                        this.applyCurrentThemeColors();
                        this.fetchThemesFromGo();
                    }
                });
            });

        },
        showError(text) {
            this.errDialogText = text;
            this.$refs.errDialog.showDialog();
        },
        createNewTheme() {
            this.$emit("createNewTheme");
        },
        fetchThemesFromGo() {
            CurrentThemeColors().then((theme) => {
                this.chosenTheme = theme;
            });
            GetAllThemesValues().then((themes) => {
                this.avaliableThemes = themes;
            });
        }
    },
    computed: {
        chosenThemeColorFields() {
            return Object.fromEntries(
                Object.entries(this.chosenTheme).filter(([key]) => key.endsWith("Color"))
            );
        },
    },
    created() {
        this.fetchThemesFromGo();
    },
    components:
    {
        ColorLabel,
        ThemeLabel,
        ErrorDialog
    },
    provide() {
        return {
            cancelColorsEditing: this.cancelAllEditing,
            updateTheme: this.updateTheme,
            showError: this.showError
        }
    },
}
</script>
<style scoped>
.visual-settings-container {
    height: calc(98% - 10px);
    width: 100%;
    display: grid;
    grid-template-columns: calc(35vw + 80px + 3vh) 1fr;
    gap: 6px;
}

.chosen-theme-zone {
    background-color: var(--back);
    height: 97%;
    width: 100%;
    display: grid;
    grid-template-rows: auto 1fr;
    padding-left: calc(6% - 5px);
}

.chosen-theme-name {
    margin-top: calc(5vh + 2px);
    margin-bottom: calc(2vh + 8px);
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(2.1vh + 6px + 0.5vw);
}

.color-labels-container {
    align-self: center;
    height: calc(65% + 100px);
    width: calc(67% + 70px);
    background-color: #f5f5f5;
    border-radius: 2%;
    display: grid;
    grid-template-rows: repeat(11, 1fr);
    padding: 2.5%;
    overflow-y: auto;
}

.avaliable-themes-zone {
    height: calc(100% - 40px - 2vw);
    width: 96%;
    padding-left: 2%;
    display: grid;
    grid-template-rows: auto 1fr calc(25px + 1.6vh + 0.5vw);

}

.avaliable-themes-label {
    width: 98%;
    padding-right: calc(0.25vw + 0.2vh + 5px);
    color: var(--front);
    font-family: 'Figtree';
    margin-top: calc(2vh + 6px);
    margin-bottom: 4px;
    font-size: calc(2vh + 9px);
    text-align: center;
    border-bottom: 1px var(--back-2) solid;

}

.themes-labels-container {
    max-height: 100%;
    overflow-y: auto;
}
.new-theme-button {
    height: 100%;
    justify-self: center;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: var(--bright-2);
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(0.55vh + 0.55vw + 11px);
    border: 1px solid transparent;
    border-radius: calc(3px + 0.09vw + 0.1vh);
    padding: calc(0.5% + 3px);
    cursor: pointer;
    gap: 3px;
    transition: 0.08s;
}

.new-theme-button:hover {
    background-color: var(--bright-3);
}

.brush-icon {
    height: calc(0.6vh + 0.6vw + 12px);
    aspect-ratio: 1/1;
}

.brush-icon path {
    stroke: var(--front);
    stroke-width: 2;
}
</style>