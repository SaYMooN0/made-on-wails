<template>
    <form class="input-container" id="newProjectForm" @submit.prevent="creationFormOnsubmit">
        <label class="new-project-header">Set up new project</label>
        <p class="new-line">
            <label class="new-project-label">Project path: </label>
        </p>
        <p class="new-line">
            <input class="new-project-input" type="text" id="path-input" required autocomplete="off" v-model="folderPath"
                :name="currentDate + 'FolderPath'">
            <button id="openFileExplorerBtn" class="open-file-explorer-btn non-selectable" type="button"
                @click="openFolderSelectionDialog">...</button>
        </p>
        <p class="new-line">
            <label class="new-project-label">Name: </label>
        </p>
        <p class="new-line">
            <input class="new-project-input" type="text" required autocomplete="off" v-model="name"
                :name="currentDate + 'Name'">
        </p>
        <p class="new-line">
            <label class="new-project-label">Version: </label>
        </p>
        <p class="new-line">
            <input class="new-project-input" type="text" required autocomplete="off" v-model="version"
                :name="currentDate + 'Version'">
        </p>
        <p class="new-line">
            <label class="new-project-label">Mod loader: </label>
        </p>
        <p class="new-line">
            <select v-model="modLoader" class="loader-select" required id="loader-input">
                <option v-for="loader in loaders" :key="loader" :value="loader">{{ loader }}</option>
            </select>
        </p>
        <button class="create-btn" type="submit">Create</button>

    </form>
    <a class="cancel-btn no-underline" @click="cancelCreation">Cancel</a>
    <ErrDialogWithGithub ref="errDialogWithGithub" :errorText="`${errDialogWithGithubText}`"> </ErrDialogWithGithub>
    <ErrorDialog ref="errDialog" :errorText="`${errDialogText}`"></ErrorDialog>
</template>

<script>
import { ChooseFolderForNewProject, GetInformationToFillCreationForm, AnyMadeProjectFilesInFolder, CreateProject, SetCurrentProjectByFolder } from "../../../wailsjs/go/projectrelated/ProjectManager";
import ErrDialogWithGithub from '../modalDialogs/ErrDialogWithGithub.vue';
import ErrorDialog from '../modalDialogs/ErrorDialog.vue';

const kjsForgeVersions = ["1.16.1", "1.16.2", "1.16.3", "1.16.4", "1.16.5", "1.18", "1.18.1", "1.18.1", "1.18.2", "1.19", "1.19.2"];
const kjsFabricVersions = ["1.18.2", "1.19", "1.19.2"];
const forbiddenNameChars = /[<>:"/\\|?*\x00-\x1F]/;
export default {
    components: {
        ErrDialogWithGithub,
        ErrorDialog
    },
    data() {
        return {
            folderPath: '',
            name: '',
            version: '',
            modLoader: 'Forge',
            loaders: ['Forge', "Fabric"],
            errDialogWithGithubText: '',
            errDialogText: '',
        };
    },
    methods: {
        openFolderSelectionDialog() {
            ChooseFolderForNewProject().then((folder) => {
                if (folder != '') {
                    AnyMadeProjectFilesInFolder(folder).then((anyProjectInFolder) => {
                        if (anyProjectInFolder) {
                            this.showNotification("!Warning! There is already a .madeProject file in this folder. You can open it by going to the starting actions page,\
                             or if you want to create a project from scratch, please delete the previous project in this folder", true);
                        }
                    });
                    GetInformationToFillCreationForm(folder).then((inf) => {
                        this.folderPath = inf.FolderPath;
                        if (inf.Version != '') {
                            this.fillForm(inf);
                        }
                        else {
                            this.showNotification('!Warning! The folder with the modpack that you are trying to open is most likely created not with CurseForge.\
                        To avoid errors in Made and the modpack itself, it is recommended to recreate your modpack using CurseForge',true);
                        }

                    });
                }
            });
        },
        fillForm(inf) {
            this.name = inf.Name;
            this.version = inf.Version;
            this.modLoader = this.loaders[inf.ModLoader];
        },
        creationFormOnsubmit() {
            if (forbiddenNameChars.test(this.name)) {
                this.errDialogText = `You can't have <, >, :, ", /, \\, |, ?, * in your project name`;
                this.$refs.errDialog.showDialog();
                return;
            }
            AnyMadeProjectFilesInFolder(this.folderPath).then((anyProjectInFolder) => {
                if (anyProjectInFolder) {
                    this.errDialogText = `Made failed to create file in folder with path:\n${folderPath}\nPlease sure there that the folder with this path exists and there is no .madeProject files. If such a file exists, you can open it by going to the starting actions page, or if you want to create a project from scratch, please delete the previous project in this folder`;
                    this.$refs.errDialog.showDialog();
                    return;
                }
                if (this.isVersionSupported()) {
                    CreateProject(this.name, this.folderPath, this.version, this.loaders.indexOf(this.modLoader)).then((creationResult) => {
                        if (creationResult === "") {
                            SetCurrentProjectByFolder(this.folderPath).then(() => {
                                this.$emit('goToProjectPage');
                            })
                        }
                        else {
                            this.errDialogWithGithubText = `An error \n${creationResult}\n occurred while creating the project. Please report it to the github issues page`;
                            this.$refs.errDialogWithGithub.showDialog();
                        }
                    })
                } else {
                    this.errDialogWithGithubText = `You are trying to create a project to work with ${modLoader} for ${version} minecraft version, but unfortunately kubejs does not support this version, therefore Made also cannot work with this version. If you are sure that kubejs supports ${version} ${modLoader} let us know on the issues page on github.`;
                    this.$refs.errDialogWithGithub.showDialog();

                }

            });
        },
        cancelCreation() {
            this.$emit('goBack');
        },
        isVersionSupported() {
            if (this.modLoader == "Forge") {
                if (kjsForgeVersions.includes(this.version)) {
                    return true;
                }
            }
            else if (this.modLoader == "Fabric") {
                if (kjsFabricVersions.includes(this.version)) {
                    return true;
                }
            }
            return false;
        },
    },
    inject: ['showNotification'],
};
</script>

<style scoped>
.input-container {
    margin-top: calc(7vh - 0.6vw);
    margin-left: 5vw;
}

.new-project-header {
    display: inline-block;
    color: var(--front);
    font-family: 'Figtree';
    letter-spacing: 1px;
    font-size: calc(2.4vh + 1.1vw + 12px);
    font-weight: 500;
    margin-bottom: calc(10px + 1%);
}

.new-line {
    margin-top: calc(2.8vh - 3px);
    align-items: center;
    display: flex;
}

.new-project-label {
    display: inline-block;
    color: var(--front);
    font-family: 'Figtree';
    letter-spacing: 1px;
    font-size: calc(0.68vh + 0.38vw + 13px);
    font-weight: 200;
}

.new-project-input {
    height: calc(2vh + 6px + 0.12vw);
    width: calc(38vw + 260px + 5vh);
    background-color: var(--back-2);
    border: 1px solid transparent;
    border-radius: calc(1px + 0.04vw + 0.1vh);
    margin-top: calc(-0.5vh - 10px);
    font-family: 'Figtree';
    font-size: calc(0.85vh + 0.3vw + 6px);
    font-weight: 300;
    color: var(--front);
}

input:focus {
    outline: none;
    background-color: var(--back-3);
}

.open-file-explorer-btn {
    height: calc(1.65vh + 12px + 0.1vw);
    width: calc(16px + 0.7vw + 1.3vh);
    align-self: center;
    background-color: var(--back-2);
    border: 1px solid transparent;
    border-radius: calc(1px + 0.06vw + 0.12vh);
    margin-left: calc(5px + 0.5vw);
    margin-top: calc(-0.6vh - 10px);
    color: var(--front);
    font-size: calc(0.53vh + 0.72vw + 11px);
    font-weight: 900;
    letter-spacing: 2px;
    font-family: 'Figtree';
    padding-left: 9px;
    display: flex;
    justify-content: center;
    align-items: center;
    padding-bottom: calc(0.14vh + 0.11vw + 3px);
}

.open-file-explorer-btn:hover {
    background-color: var(--back-3);
    cursor: pointer;
}

.create-btn {
    position: absolute;
    bottom: calc(10px + 2vh);
    right: calc(13px + 1vw + 0.8vh);
    height: calc(18px + 1.8vh + 0.3vw);
    width: calc(36px + 1vh + 3.9vw);
    background-color: var(--back-3);
    border: none;
    font-family: 'Figtree';
    font-size: calc(0.33vh + 0.48vw + 8px);
    font-weight: 200;
    color: var(--front);
    border: 1px solid transparent;
    border-radius: calc(1px + 0.04vw + 0.1vh)
}

.create-btn:hover {
    background-color: var(--bright-2);
    cursor: pointer;
}

.cancel-btn {
    position: absolute;
    bottom: calc(10px + 2vh);
    right: calc(1vh + 57px + 6vw);
    height: calc(18px + 1.8vh + 0.3vw);
    width: calc(36px + 1vh + 3.9vw);
    background-color: var(--back-3);
    display: flex;
    justify-content: center;
    align-items: center;
    font-family: 'Figtree';
    font-size: calc(0.33vh + 0.48vw + 8px);
    font-weight: 200;
    color: var(--front);
    border: 1px solid transparent;
    border-radius: calc(1px + 0.04vw + 0.1vh)
}

.cancel-btn:hover {
    background-color: var(--bright-2);
    cursor: pointer;
}

.loader-select {
    color: var(--front);
    background-color: var(--back-2);
    height: calc(0.8vh + 19px + 0.4vw);
    width: calc(30px + 6vw + 1vh);
    font-size: calc(0.48vh + 0.52vw + 8px);
    font-family: 'Figtree';
    margin-top: calc(-0.8vh - 6px);
    outline: none;
    border: 1px solid transparent;
    border-radius: calc(1px + 0.04vw + 0.1vh);
}
</style>