<template>
    <div class="visual-settings-container">
        <div class="one-theme-zone">
            <div v-for="(value, key) in chosenThemeColorFields" :key="key">
                <label>{{ key }}: </label>{{ value }}
            </div>
        </div>
        <div class="avaliable-themes-zone">

        </div>
    </div>
</template>
<script>


import { CurrentThemeColors, GetAllThemesValues } from "./../../../../wailsjs/go/themerelated/ThemeCollection";

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
    }
}
</script>
<style scoped>
.visual-settings-container {
    background-color: red;
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: 1.2fr 1fr;
}

.one-theme-zone {
    background-color: cadetblue;
    height: 100%;
    width: 100%;
}

.avaliable-themes-zone {
    background-color: violet;
    height: 100%;
    width: 100%;
}
</style>