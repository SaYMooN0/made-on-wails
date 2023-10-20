<template>
    <label id="project-header" class="project-header">Your projects</label>

    <div class="projects-container" id="projects-container">
        <ProjectItem v-for="project in  projects" :key="project.id" :project="project" @goToProjectPage="$emit('goToProjectPage')"/>
    </div>

    <div class="action-btns-container">
        <p class="work-starting-text">Start working</p>
        <ActionButton link="/projectCreation" name="Create new project" @createNewProject="$emit('createNewProject')"
            info="Open already existing mod pack folder and start using Made in it">
            <svg class="action-icon" viewBox="0 0 24 24" fill="none">
                <path
                    d="M20 10.5V6.8C20 5.11984 20 4.27976 19.673 3.63803C19.3854 3.07354 18.9265 2.6146 18.362 2.32698C17.7202 2 16.8802 2 15.2 2H8.8C7.11984 2 6.27976 2 5.63803 2.32698C5.07354 2.6146 4.6146 3.07354 4.32698 3.63803C4 4.27976 4 5.11984 4 6.8V17.2C4 18.8802 4 19.7202 4.32698 20.362C4.6146 20.9265 5.07354 21.3854 5.63803 21.673C6.27976 22 7.11984 22 8.8 22H12M14 11H8M10 15H8M16 7H8M18 21V15M15 18H21"
                    stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
        </ActionButton>
        <ActionButton name="Open project" info="Open an already existing .madeProject file and modify it">
            <svg class="action-icon" viewBox="0 0 24 24" fill="none">
                <path
                    d="M14 11H8M10 15H8M16 7H8M20 10.5V6.8C20 5.11984 20 4.27976 19.673 3.63803C19.3854 3.07354 18.9265 2.6146 18.362 2.32698C17.7202 2 16.8802 2 15.2 2H8.8C7.11984 2 6.27976 2 5.63803 2.32698C5.07354 2.6146 4.6146 3.07354 4.32698 3.63803C4 4.27976 4 5.11984 4 6.8V17.2C4 18.8802 4 19.7202 4.32698 20.362C4.6146 20.9265 5.07354 21.3854 5.63803 21.673C6.27976 22 7.11984 22 8.8 22H11.5M22 22L20.5 20.5M21.5 18C21.5 19.933 19.933 21.5 18 21.5C16.067 21.5 14.5 19.933 14.5 18C14.5 16.067 16.067 14.5 18 14.5C19.933 14.5 21.5 16.067 21.5 18Z"
                    stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
        </ActionButton>
        <ActionButton name="Use without mod pack" info="Use Made functionality without mod pack">
            <svg class="action-icon" viewBox="0 0 24 24" fill="none">
                <path
                    d="M20 11.9412V6.8C20 5.11984 20 4.27976 19.673 3.63803C19.3854 3.07354 18.9265 2.6146 18.362 2.32698C17.7202 2 16.8802 2 15.2 2H8.8C7.11984 2 6.27976 2 5.63803 2.32698C5.07354 2.6146 4.6146 3.07354 4.32698 3.63803C4 4.27976 4 5.11984 4 6.8V17.2C4 18.8802 4 19.7202 4.32698 20.362C4.6146 20.9265 5.07354 21.3854 5.63803 21.673C6.27976 22 7.11984 22 8.8 22H14M14 11H8M10 15H8M16 7H8M15 17H21"
                    stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
        </ActionButton>
    </div>
    <div class="right-bottom-buttons-container">
        <LinkButton iconColor="fill" @click="openGitHub">
            <svg viewBox="-2.5 0 19 19">
                <path
                    d="M9.464 17.178a4.506 4.506 0 0 1-2.013.317 4.29 4.29 0 0 1-2.007-.317.746.746 0 0 1-.277-.587c0-.22-.008-.798-.012-1.567-2.564.557-3.105-1.236-3.105-1.236a2.44 2.44 0 0 0-1.024-1.348c-.836-.572.063-.56.063-.56a1.937 1.937 0 0 1 1.412.95 1.962 1.962 0 0 0 2.682.765 1.971 1.971 0 0 1 .586-1.233c-2.046-.232-4.198-1.023-4.198-4.554a3.566 3.566 0 0 1 .948-2.474 3.313 3.313 0 0 1 .091-2.438s.773-.248 2.534.945a8.727 8.727 0 0 1 4.615 0c1.76-1.193 2.532-.945 2.532-.945a3.31 3.31 0 0 1 .092 2.438 3.562 3.562 0 0 1 .947 2.474c0 3.54-2.155 4.32-4.208 4.548a2.195 2.195 0 0 1 .625 1.706c0 1.232-.011 2.227-.011 2.529a.694.694 0 0 1-.272.587z" />
            </svg>
        </LinkButton>
        <LinkButton iconColor="fill" @click="$emit('changeToInfo')">
            <svg viewBox="0 0 24 24">
                <path fill-rule="evenodd" clip-rule="evenodd"
                    d="M12 7C12.8284 7 13.5 6.32843 13.5 5.5C13.5 4.67157 12.8284 4 12 4C11.1716 4 10.5 4.67157 10.5 5.5C10.5 6.32843 11.1716 7 12 7ZM11 9C10.4477 9 10 9.44772 10 10C10 10.5523 10.4477 11 11 11V19C11 19.5523 11.4477 20 12 20C12.5523 20 13 19.5523 13 19V10C13 9.44772 12.5523 9 12 9H11Z" />
            </svg>
        </LinkButton>
        <LinkButton iconColor="stroke" @click="$emit('changeToSettings')">
            <svg viewBox="-0.6 -1 25 25" fill="none">
                <circle cx="12" cy="12" r="3" stroke-width="1.5" />
                <path
                    d="M13.7654 2.15224C13.3978 2 12.9319 2 12 2C11.0681 2 10.6022 2 10.2346 2.15224C9.74457 2.35523 9.35522 2.74458 9.15223 3.23463C9.05957 3.45834 9.0233 3.7185 9.00911 4.09799C8.98826 4.65568 8.70226 5.17189 8.21894 5.45093C7.73564 5.72996 7.14559 5.71954 6.65219 5.45876C6.31645 5.2813 6.07301 5.18262 5.83294 5.15102C5.30704 5.08178 4.77518 5.22429 4.35436 5.5472C4.03874 5.78938 3.80577 6.1929 3.33983 6.99993C2.87389 7.80697 2.64092 8.21048 2.58899 8.60491C2.51976 9.1308 2.66227 9.66266 2.98518 10.0835C3.13256 10.2756 3.3397 10.437 3.66119 10.639C4.1338 10.936 4.43789 11.4419 4.43786 12C4.43783 12.5581 4.13375 13.0639 3.66118 13.3608C3.33965 13.5629 3.13248 13.7244 2.98508 13.9165C2.66217 14.3373 2.51966 14.8691 2.5889 15.395C2.64082 15.7894 2.87379 16.193 3.33973 17C3.80568 17.807 4.03865 18.2106 4.35426 18.4527C4.77508 18.7756 5.30694 18.9181 5.83284 18.8489C6.07289 18.8173 6.31632 18.7186 6.65204 18.5412C7.14547 18.2804 7.73556 18.27 8.2189 18.549C8.70224 18.8281 8.98826 19.3443 9.00911 19.9021C9.02331 20.2815 9.05957 20.5417 9.15223 20.7654C9.35522 21.2554 9.74457 21.6448 10.2346 21.8478C10.6022 22 11.0681 22 12 22C12.9319 22 13.3978 22 13.7654 21.8478C14.2554 21.6448 14.6448 21.2554 14.8477 20.7654C14.9404 20.5417 14.9767 20.2815 14.9909 19.902C15.0117 19.3443 15.2977 18.8281 15.781 18.549C16.2643 18.2699 16.8544 18.2804 17.3479 18.5412C17.6836 18.7186 17.927 18.8172 18.167 18.8488C18.6929 18.9181 19.2248 18.7756 19.6456 18.4527C19.9612 18.2105 20.1942 17.807 20.6601 16.9999C21.1261 16.1929 21.3591 15.7894 21.411 15.395C21.4802 14.8691 21.3377 14.3372 21.0148 13.9164C20.8674 13.7243 20.6602 13.5628 20.3387 13.3608C19.8662 13.0639 19.5621 12.558 19.5621 11.9999C19.5621 11.4418 19.8662 10.9361 20.3387 10.6392C20.6603 10.4371 20.8675 10.2757 21.0149 10.0835C21.3378 9.66273 21.4803 9.13087 21.4111 8.60497C21.3592 8.21055 21.1262 7.80703 20.6602 7C20.1943 6.19297 19.9613 5.78945 19.6457 5.54727C19.2249 5.22436 18.693 5.08185 18.1671 5.15109C17.9271 5.18269 17.6837 5.28136 17.3479 5.4588C16.8545 5.71959 16.2644 5.73002 15.7811 5.45096C15.2977 5.17191 15.0117 4.65566 14.9909 4.09794C14.9767 3.71848 14.9404 3.45833 14.8477 3.23463C14.6448 2.74458 14.2554 2.35523 13.7654 2.15224Z"
                    stroke-width="1.5" />
            </svg>
        </LinkButton>
    </div>
</template>
  
<script>
import { GetProjects } from "../../../wailsjs/go/projectrelated/ProjectManager";

import ProjectItem from '../ProjectItem.vue';
import ActionButton from '../ActionButton.vue';
import LinkButton from '../LinkButton.vue';


export default {
    components: {
        ActionButton,
        ProjectItem,
        LinkButton
    },
    data() {
        return {
            pinnedProjects: [],
            projects: []
        };
    },
    created() {
        this.fetchPinnedProjects();
        this.fetchProjects();
    },
    methods: {
        projectItemClicked(FolderPath) {
            // const selectedProject = this.projects.find(i => i.PathToFolder === FolderPath);
            // this.$router.push("/projectPage");
        },
        openProjectItemClicked() {
            // const pathToProject = this.chooseProject();
            // if (pathToProject) {
            //     const selectedProject = this.projects.find(i => i.FullPath === pathToProject);
            //     this.$router.push("/projectPage");
            // }
        },
        fetchPinnedProjects() {
            // GetProjects().then((projectsToFetch) => {
            //     this.projects = projectsToFetch;
            // });
        },
        fetchProjects() {
            GetProjects().then((projectsToFetch) => {
                this.projects = projectsToFetch;
            });
        },

        chooseProject() {
            // Logic for choosing a project
            return null;
        },
        openGitHub() {
            window.open('https://github.com/SaYMooN0/made-on-wails');
        }
    }
};
</script>
<style scoped>
.project-header {
    position: absolute;
    top: calc(3.8vh + 12px);
    left: calc(2vw + 12px);
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(2.3vh + 2vw + 8px);
    font-weight: 600;
    letter-spacing: 2px;
}

.projects-container {
    position: absolute;
    top: calc(9.5vh + 26px + 3vw);
    left: calc(2vw + 10px);
    width: calc(100% - 32px - 32vw);
    height: 65%;
    overflow-y: auto;
}

.projects-container::-webkit-scrollbar {
    width: calc(0.25vw + 0.2vh + 4px);

}

.projects-container::-webkit-scrollbar-track {
    background: var(--back-3);
    border-radius: calc(0.1vw + 0.04vh + 2px);
}

.projects-container::-webkit-scrollbar-thumb {
    background: var(--front-2);
    border-radius: calc(0.1vw + 0.04vh + 2px);
}

.projects-container::-webkit-scrollbar-thumb:hover {
    background: var(--front-3);
}



.corrupted-sign-dialog {
    position: absolute;
    top: 0;
    bottom: 10%;
    width: 46%;
    background-color: var(--back-2);
    border-radius: calc(0.45vh + 0.45vw);
    border-width: calc(0.05vh + 0.05vw + 4px);
    border-color: var(--warning-bright);
}

dialog:focus {
    outline: none;
}

.corrupted-sign-dialog-container {
    display: grid;
    grid-template-rows: 8fr 1fr;
    grid-template-columns: 9fr 2.2fr 1fr 2.2fr 1fr;
    grid-template-areas:
        "txt txt txt txt txt"
        ".  btn-cncl .  btn-dlt . ";
}

.csd-txt {
    grid-area: txt;
    font-size: calc(1.1vh + 1.1vw);
    color: var(--front);
    font-family: 'Bahnschrift';
    font-weight: 400;
    text-align: justify;
}

.csd-btn {
    background-color: var(--bright);
    width: 100%;
    grid-area: btn-cncl;
    display: flex;
    justify-content: center;
    align-items: center;
    font-family: 'Figtree';
    font-size: calc(0.75vh + 0.52vw + 2px);
    letter-spacing: 0.1vw;
    color: var(--front);
}

.csd-btn:hover {
    background-color: var(--bright-3);
    cursor: pointer;
}

.csd-btn-cncl {
    grid-area: btn-cncl;
}

.csd-btn-dlt {
    grid-area: btn-dlt;
}

.action-btns-container {
    position: absolute;
    width: calc(76px + 14vw + 3vh);
    min-width: 200px;
    height: calc(18vh + 8vw + 120px);
    right: calc(5vw - 15px);
    top: 12%;
    display: grid;
    grid-template-rows: 1fr 3fr 3fr 3fr;
    grid-gap: calc(5px + 1vh);

}

.work-starting-text {
    height: 100%;
    font-family: 'Figtree';
    display: flex;
    align-items: center;
    color: var(--front);
    font-size: calc(1.25vh + 1.25vw + 4px);
    letter-spacing: 2px;
    margin-top: 0;
}

.right-bottom-buttons-container {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    gap: calc(4px + 4%);
    position: absolute;
    right: calc(5vw - 10px);
    bottom: 14px;
}
</style>
  