<template>
    <div class="visual-settings-container">
        <div class="chosen-theme-zone">
            <label class="chosen-theme-name">Theme: {{ chosenTheme.Name }}</label>
            <div class="color-labels-container">
                <ColorLabel v-for="(value, key) in chosenThemeColorFields" :key="key" :label="key" :color="value" />
            </div>

        </div>
        <div class="avaliable-themes-zone">
            <p class="avaliable-themes-label">Available themes:</p>
            <div class="themes-labels-container">
                <ThemeLabel v-for="theme in avaliableThemes" :theme="theme" />
            </div>
            <button class="new-theme-button">Create new theme</button>
        </div>
    </div>
</template>
<script>


import { CurrentThemeColors, GetAllThemesValues } from "./../../../../wailsjs/go/themerelated/ThemeCollection";
import ColorLabel from "../tabsComponents/ColorLabel.vue";
import ThemeLabel from "../tabsComponents/ThemeLabel.vue";

export default {
    data() {
        return {
            avaliableThemes: [],
            chosenTheme: {}
        };
    },
    methods: {
        changeTheme() {
            CurrentThemeColors().then((theme) => {
                document.documentElement.style.setProperty('--back', theme.MainBackColor);
                document.documentElement.style.setProperty('--back-2', theme.SecondBackColor);
                document.documentElement.style.setProperty('--back-3', theme.ThirdBackColor);
                document.documentElement.style.setProperty('--front', theme.MainFrontColor);
                document.documentElement.style.setProperty('--front-2', theme.SecondFrontColor);
                document.documentElement.style.setProperty('--front-3', theme.ThirdFrontColor);
                document.documentElement.style.setProperty('--bright', theme.MainBrightColor);
                document.documentElement.style.setProperty('--bright-2', theme.SecondBrightColor);
                document.documentElement.style.setProperty('--bright-3', theme.ThirdBrightColor);
                document.documentElement.style.setProperty('--warning-main', theme.WarningMainColor);
                document.documentElement.style.setProperty('--warning-bright', theme.WarningBrightColor);
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
        CurrentThemeColors().then((theme) => {
            this.chosenTheme = theme;
        });
        GetAllThemesValues().then((themes) => {
            this.avaliableThemes = themes;
        });
    },
    components:
    {
        ColorLabel,
        ThemeLabel
    }
}
</script>
<style scoped>
.visual-settings-container {
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: calc(35vw + 80px + 3vh) 1fr;
    gap: 6px;
}

.chosen-theme-zone {
    background-color: var(--back);
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-rows: calc(24vh - 40px) 1fr;
    padding-left: calc(6% - 5px);
}

.chosen-theme-name {
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(1.2vh + 1.55vw + 11px);
    padding-left: calc(4px + 1vw - 1vh);
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
}

.color-labels-container {
    width: calc(68% + 64px);
    height: 100%;
    max-height: calc(50px + 54vh);
    align-self: center;
    background-color: #f5f5f5;
    border-radius: 2%;
    display: grid;
    grid-template-rows: repeat(11, 1fr);
    padding: 4%;
    overflow-y: auto;
}

.avaliable-themes-zone {
    height: 100%;
    width: 94%;
    padding-left: 3%;
    padding-right: 3%;
}

.avaliable-themes-label {
    width: 100%;
    color: var(--front);
    font-family: 'Figtree';
    margin-top: calc(1vh + 0.5vw + 10px);
    margin-bottom: 4px;
    font-size: calc(0.9vh + 1.2vw + 9px);
    text-align: center;
    border-bottom: 2px var(--back-2) solid;
}
</style>