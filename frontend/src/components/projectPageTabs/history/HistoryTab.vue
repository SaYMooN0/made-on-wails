<template>
    <div class='history-page-top-container'>
        <label class='project-history-label'>Project history</label>
        <div class='history-search-container' @click="focusSearch">
            <svg class='history-search-icon' viewBox='0 0 16 16'>
                <rect width='16' height='16' id='icon-bound' fill='none' />
                <path
                    d='M15.708,13.587l-3.675-3.675C12.646,8.92,13,7.751,13,6.5C13,2.91,10.09,0,6.5,0S0,2.91,0,6.5S2.91,13,6.5,13 c1.251,0,2.42-0.354,3.412-0.967l3.675,3.675c0.389,0.389,1.025,0.389,1.414,0l0.707-0.707 C16.097,14.612,16.097,13.976,15.708,13.587z M3.318,9.682C2.468,8.832,2,7.702,2,6.5s0.468-2.332,1.318-3.182S5.298,2,6.5,2 s2.332,0.468,3.182,1.318C10.532,4.168,11,5.298,11,6.5s-0.468,2.332-1.318,3.182C8.832,10.532,7.702,11,6.5,11 S4.168,10.532,3.318,9.682z' />
            </svg>
            <input type='text' class='history-search-input' @input='filterHistoryItems' />
        </div>
    </div>
    <div class='history-items-container scrollbar'>
        <div v-for="item in historyItems" :key="item.ActionID">
            <HistoryItem :historyItem="item" @open="openHistoryItem(historyItem)" @delete="deleteHistoryItem(item)" />
        </div>
    </div>

    <EmptyDialog class='histoty-item-deleting-dialog' ref="deletingDialog">
        <p class='history-item-deletig-dialog-main-label'>Are you sure you want to delete this action?</p>
        <p class='dont-show-label'>Don't show this message in this project anymore
            <DefCheckBox v-model="doNotShowWarningWhenDeletingAction" class="is-shapeless-checkbox" />

        </p>

        <div class="dialog-buttons">
            <button @click="closeDialog">Cancel</button>
            <button @click="confirmDeleting(true)">Delete</button>
        </div>
    </EmptyDialog>
</template>
<script>
import {
    CurrentProjectHistory, DoShowWarningWhenDeletingActionForCurrentProjectHistory, CurrentProjectChangeAction,
    CurrentProjectDeleteAction,
    SetDoShowWarningWhenDeletingActionForCurrentProjectHistory
} from "../../../../wailsjs/go/projectrelated/ProjectManager";

import EmptyDialog from "../../modalDialogs/EmptyDialog.vue"
import DefCheckBox from "../../default/DefCheckBox.vue"
import HistoryItem from './HistoryItem.vue';
export default {
    components: {
        EmptyDialog,
        DefCheckBox,
        HistoryItem
    },
    data() {
        return {
            historyItems: [],
            doNotShowWarningWhenDeletingAction: false,
            actionToDelete: {},
        };
    },
    methods: {
        focusSearch() {
        },
        filterHistoryItems() {
        },
        openHistoryItem(historyItem) {
        },
        deleteHistoryItem(historyItem) {
            this.actionToDelete = historyItem;
            DoShowWarningWhenDeletingActionForCurrentProjectHistory().then((doShow) => {
                if (doShow) {
                    this.$refs.deletingDialog.showDialog();
                }
                else {
                    this.confirmDeleting();
                }
            });
        },
        closeDialog() {
            this.$refs.deletingDialog.closeDialog();
        },
        confirmDeleting(changeDoNotShowSettings = false) {
            if (changeDoNotShowSettings) {
                SetDoShowWarningWhenDeletingActionForCurrentProjectHistory(false);
            }
            CurrentProjectDeleteAction(this.actionToDelete.ActionID, this.actionToDelete.FilePath);

            const index = this.historyItems.findIndex(obj => obj === this.actionToDelete);
            if (index > -1) {  this.historyItems.splice(index, 1);  }
        },
        fetchHistoryItems() {

            CurrentProjectHistory().then((history) => {
                this.historyItems = history;
            });
        }
    },
    mounted() {
        this.fetchHistoryItems();
    }
}
</script>

<style scoped>
.history-page-top-container {
    position: absolute;
    top: calc(30px + 4%);
    left: 3%;
    width: 96%;
    height: calc(20px + 5%);
    z-index: 100;
    display: grid;
    grid-template-columns: 1fr calc(150px - 7vh + 19vw);
    align-items: center;
}

.project-history-label {
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(1.4vh + 1.4vw + 11px);
}

.history-search-container {
    justify-self: right;
    width: 90%;
    height: calc(16px + 1.6vh + 1.2vw);
    background-color: var(--back-2);
    border-radius: calc(100vw + 100vh);
    display: grid;
    align-items: center;
    grid-template-columns: calc(11px + 0.7vh + 0.45vw) 1fr;
    gap: calc(2px + 0.2vw + 0.2vh);
    padding-left: calc(8px + 0.5vw);
}

.history-search-input {
    width: 95%;
    height: 100%;
    border: none;
    background: none;
    font-size: calc(0.5vh + 0.6vw + 8px);
    color: var(--front);
    outline: none;
    font-family: Bahnschrift;
}

.history-search-icon {
    width: 100%;
    aspect-ratio: 1/1;
    fill: var(--front);
    pointer-events: none;
}


.history-items-container {
    position: absolute;
    width: 98%;
    height: calc(90% - 60px - 1vw);
    overflow-y: auto;
    bottom: 2%;
    left: 1%;
}


.histoty-item-deleting-dialog {
    border-color: var(--warning-bright);
    padding-top: 0;
    padding-left: 1%;
    padding-right: 1%;
}

.history-item-deletig-dialog-main-label {
    font-family: 'Figtree';
    color: var(--front);
    font-size: calc(0.4vh + 0.65vw + 10px);
    font-weight: 200;
    width: 100%;
    text-align: center;
}

.dont-show-label {
    font-family: 'Figtree';
    color: var(--front);
    font-size: calc(0.4vh + 0.65vw + 8px);
    font-weight: 200;
    width: 100%;
    text-align: center;
}

.dialog-buttons {
    display: inline-flex;
    width: 100%;
    justify-content: right;
    gap: calc(4px + 0.5vw);
}
</style>
