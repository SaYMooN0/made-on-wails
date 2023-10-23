<template>
    <h1 class="welcome-header">Welcome to Made</h1>
    <h5 class="welcome-reminder">Make sure you have downloaded KubeJS of the required version before you start working. If
        not yet, you can download KubeJS <a href="https://www.curseforge.com/minecraft/mc-mods/kubejs/files"
            target="_blank">here</a></h5>
</template>
<style scoped>
.welcome-header {
    position: absolute;
    color: var(--bright);
    font-family: 'Figtree';
    letter-spacing: 2px;
    font-size: calc(1.7vh + 1.4vw + 10px);
    margin-left: calc(20px + 2vw);
    margin-top: calc(10px + 2vh + 0.7vw);
    width: 70%;
    display: block;
}

.welcome-reminder {
    position: absolute;
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(0.7vh + 0.9vw + 7px);
    margin-left: calc(20px + 2vw);
    margin-right: calc(20px + 2vw);
    margin-top: calc(11vh + 18px + 1vw);
    text-align: justify;
}

.welcome-reminder a {
    color: var(--front-2);
}

.welcome-reminder a:hover {
    color: var(--front-3);
}
</style>