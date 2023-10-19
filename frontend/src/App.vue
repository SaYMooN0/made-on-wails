<template>
    <div>
        <div v-if="currentPage === 'startingActions'">
            <StartingActionsPage @createNewProject="currentPage = 'projectCreation'"
                @changeToInfo="currentPage = 'madeInfo'" @changeToSettings="currentPage = 'settings'" />
        </div>
        <div v-if="currentPage === 'settings'">
            <SettigsPage @goBack="currentPage = 'startingActions'" />
        </div>
        <div v-if="currentPage === 'projectCreation'">
            <ProjectCreationPage @goBack="currentPage = 'startingActions'" @goToProjectPage="currentPage = 'projectEditing'"/>
        </div>
        <div v-if="currentPage === 'madeInfo'">
            <MadeInfoPage @goBack="currentPage = 'startingActions'" />
        </div>
        <div v-if="currentPage === 'projectEditing'">
            <ProjectPage @goBack="currentPage = 'startingActions'"/>
        </div>
    </div>
</template>

  
<script>
import StartingActionsPage from './components/pages/StartingActionsPage.vue';
import SettigsPage from './components/pages/SettigsPage.vue';
import ProjectCreationPage from './components/pages/ProjectCreationPage.vue';
import MadeInfoPage from './components/pages/MadeInfoPage.vue';
import ProjectPage from './components/pages/ProjectPage.vue';


import { CurrentThemeColors } from "../wailsjs/go/themerelated/ThemeCollection";

export default {
    components: {
        StartingActionsPage,
        SettigsPage,
        ProjectCreationPage,
        MadeInfoPage,
        ProjectPage
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