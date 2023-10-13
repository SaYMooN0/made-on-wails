<template>
    <div>
        <label id="project-header" class="project-header">Your projects</label>

        <div class="projects-container" id="projects-container">
            <ProjectItem v-for="project in pinnedProjects" :key="project.id" :project="project" />
            <div v-if="pinnedProjects && pinnedProjects.length" class="pinned-projects-divider"></div>
            <ProjectItem v-for="project in projects" :key="project.id" :project="project" />
        </div>

        <div class="action-btns-container">
            <p class="work-starting-text">Start working</p>
            <ActionButton link="/ProjectCreationPage" name="Create new project"
                @createNewProject="$emit('createNewProject')"
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
    </div>
</template>
  
<script>
import ProjectItem from './ProjectItem.vue';
import ActionButton from './ActionButton.vue';


export default {
    components: {
        ActionButton,
        ProjectItem
    },
    data() {
        return {
            pinnedProjects: [],
            projects: []
        };
    },
    created() {
        // this.pinnedProjects = this.fetchPinnedProjects();
        // this.projects = this.fetchProjects();
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
            // Fetch your pinned projects from your services/APIs
            return [];
        },
        fetchProjects() {
            // Fetch your projects from your services/APIs
            return [];
        },
        chooseProject() {
            // Logic for choosing a project
            return null;
        }
    }
};
</script>
<style>
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
    right: calc(7vw - 26px);
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
</style>
  