<template>
    <div>
        <div v-if="currentPage === 'startingActions'">
            <StartingActionsPage />
        </div>
        <div v-if="currentPage === 'settings'">
            <SettigsPage @goBack="currentPage = 'startingActions'" />
        </div>
    </div>
</template>
  
<script>
import StartingActionsPage from './components/StartingActionsPage.vue';
import SettigsPage from './components/SettigsPage.vue';
import { CurrentTheme } from "../wailsjs/go/made/ThemeCollection";

export default {
    components: {
        StartingActionsPage,
        SettigsPage
    },
    data() {
        return {
            currentPage: 'startingActions'
        };
    },
    mounted() {
        this.bindTheme();
    },
    methods: {
        bindTheme() {
            CurrentTheme().then((theme) => {
                console.log(theme);
                document.documentElement.style.setProperty('--back', theme.mainBackColor);
                document.documentElement.style.setProperty('--back-2', theme.secondBackColor);
                document.documentElement.style.setProperty('--back-3', theme.thirdBackColor);
                document.documentElement.style.setProperty('--front', theme.mainFrontColor);
                document.documentElement.style.setProperty('--front-2', theme.secondFrontColor);
                document.documentElement.style.setProperty('--front-3', theme.thirdFrontColor);
                document.documentElement.style.setProperty('--bright', theme.mainBrightColor);
                document.documentElement.style.setProperty('--bright-2', theme.secondBrightColor);
                document.documentElement.style.setProperty('--bright-3', theme.thirdBrightColor);
                document.documentElement.style.setProperty('--warning-main', theme.warningMainColor);
                document.documentElement.style.setProperty('--warning-bright', theme.warningBrightColor);
            });
        }
    }
}
</script>
<style>
:root {
    --back: #aabb12;
    --back-2: #23bfac;
    --back-3: #cbcb11;
    --front: #232bbc;
    --front-2: #87abfa;
    --front-3: #43bbaa;
    --bright: #bb6677;
    --bright-2: #aa545b;
    --bright-3: #21bb88;
    --warning-main: #ff11ff;
    --warning-bright: #bbff43;
}

body {
    background-color: var(--back);
}
</style>