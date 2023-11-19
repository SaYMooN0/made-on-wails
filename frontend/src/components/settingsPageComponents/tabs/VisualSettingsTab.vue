<template>
    <div class="visual-settings-container">
        <div class="chosen-theme-zone">
            <label class="chosen-theme-name">Theme: {{chosenTheme.Name}}</label>
            <div class="color-labels-container">
                <ColorLabel v-for="(value, key) in chosenThemeColorFields" :key="key" :label="key" :color="value" />
            </div>

        </div>
        <div class="avaliable-themes-zone">

        </div>
    </div>
</template>
<script>


import { CurrentThemeColors, GetAllThemesValues } from "./../../../../wailsjs/go/themerelated/ThemeCollection";
import ColorLabel from "../tabsComponents/ColorLabel.vue";
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
        ColorLabel
    }
}
</script>
<style scoped>
.visual-settings-container {
    background-color: red;
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: calc(35vw + 60px) 1fr;
}

.chosen-theme-zone {
    background-color: var(--back);
    height: 84%;
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding:8% 0 10% 4%;
    gap:10%;
}
.chosen-theme-name
{
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(1.4vh + 1.4vw + 11px);
    padding-left: 8px;
}
.color-labels-container
{
    justify-self: center;
    width: 70%;
    height:calc(50% + 120px);
    padding:5vh 4% 5vh 8%;
    background-color: #f5f5f5;
    border-radius: 2%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}
.avaliable-themes-zone {
    background-color: violet;
    height: 100%;
    width: 100%;
}

</style>