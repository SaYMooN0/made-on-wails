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
    <div class='history-items-container'>
        <EmptyDialog class='histoty-item-deleting-dialog' ref="deletingDialog">
            <p class='history-item-deletig-dialog-main-label'>Are you sure you want to delete this action?</p>
            <p class='dont-show-label'>Don't show this message in this project anymore
                    <DefCheckBox v-model="isShapelessValue" class="is-shapeless-checkbox" />
           
            </p>
            
            <div class="dialog-buttons">
                <button @click="closeDialog">Cancel</button>
                <button @click="confirmDeleting(historyItem)">Delete</button>
            </div>
        </EmptyDialog>
        <div v-for="historyItem in historyItems" class='history-item' @click="openHistoryItem(historyItem)">
            <label class='history-item-type'>{{ actionTypeToString(historyItem.ActionType) }}</label>
            <svg class='history-item-delete-button' @click.stop="deleteHistoryItem(historyItem)" viewBox='0 0 24 24'
                fill='none'>
                <path stroke='#1C274C' d='M20.5001 6H3.5' stroke-width='1.5' stroke-linecap='round' />
                <path
                    d='M18.8332 8.5L18.3732 15.3991C18.1962 18.054 18.1077 19.3815 17.2427 20.1907C16.3777 21 15.0473 21 12.3865 21H11.6132C8.95235 21 7.62195 21 6.75694 20.1907C5.89194 19.3815 5.80344 18.054 5.62644 15.3991L5.1665 8.5'
                    stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
                <path d='M9.5 11L10 16' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
                <path d='M14.5 11L14 16' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
                <path
                    d='M6.5 6C6.55588 6 6.58382 6 6.60915 5.99936C7.43259 5.97849 8.15902 5.45491 8.43922 4.68032C8.44784 4.65649 8.45667 4.62999 8.47434 4.57697L8.57143 4.28571C8.65431 4.03708 8.69575 3.91276 8.75071 3.8072C8.97001 3.38607 9.37574 3.09364 9.84461 3.01877C9.96213 3 10.0932 3 10.3553 3H13.6447C13.9068 3 14.0379 3 14.1554 3.01877C14.6243 3.09364 15.03 3.38607 15.2493 3.8072C15.3043 3.91276 15.3457 4.03708 15.4286 4.28571L15.5257 4.57697C15.5433 4.62992 15.5522 4.65651 15.5608 4.68032C15.841 5.45491 16.5674 5.97849 17.3909 5.99936C17.4162 6 17.4441 6 17.5 6'
                    stroke-width='1.5' />
            </svg>

            <label class='history-item-date'>{{ historyItem.ActionID }}</label>
            <label class='history-item-input-output'>{{ historyItemLabel(historyItem) }}</label>
        </div>
    </div>
</template>
<script>
import { CurrentProjectHistory, DoShowWarningWhenDeletingActionForCurrentProjectHistory, CurrentProjectChangeAction } from "../../../../wailsjs/go/projectrelated/ProjectManager";
import { actionTypeToString, historyItemLabel } from "../../../HistoryItemManager";
import EmptyDialog from "../../modalDialogs/EmptyDialog.vue"
import DefCheckBox from "../../default/DefCheckBox.vue"

export default {
    components: {
        EmptyDialog,
        DefCheckBox
    },
    data() {
        return {
            historyItems: [],
        };
    },
    methods: {
        actionTypeToString,
        historyItemLabel,
        focusSearch() {
        },
        filterHistoryItems() {
        },
        openHistoryItem(historyItem) {
        },
        deleteHistoryItem(historyItem) {
            DoShowWarningWhenDeletingActionForCurrentProjectHistory().then((doShow) => {
                if (doShow) {
                    this.$refs.deletingDialog.showDialog();
                }
            });
        },
        closeDialog() {
            this.$refs.deletingDialog.closeDialog();
        },
        confirmDeleting(historyItem) {

        },
        fetchHistoryItems() {

            CurrentProjectHistory().then((history) => {
                console.log(history);
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

.history-items-container::-webkit-scrollbar {
    width: calc(0.2vw + 0.2vh + 5px);
}

.history-items-container::-webkit-scrollbar-track {
    background: var(--back-3);
    border-radius: calc(0.08vw + 0.02vh + 2px);
}

.history-items-container::-webkit-scrollbar-thumb {
    background: var(--front-2);
    border-radius: calc(0.08vw + 0.02vh + 2px);
}

.history-items-container::-webkit-scrollbar-thumb:hover {
    background: var(--front-3);
}

.history-item {
    min-height: 52px;
    height: calc(25px + 5.5vh + 0.9vw);
    width: 97%;
    background-color: var(--back);
    border-radius: calc(0.14vw + 0.12vh + 3px);
    cursor: pointer;
    display: grid;
    grid-template-rows: 1.2fr 1fr;
    grid-template-columns: 8fr 4fr calc(10px + 3vh + 2vw);
    grid-template-areas:
        "type date dlt"
        "text text text";
    padding-top: 1%;
    padding-left: 2%;
}


.history-item:hover {
    background-color: var(--back-2);
}


.history-item-date {
    color: var(--front-2);
    grid-area: date;
    font-family: 'Bahnschrift';
    cursor: inherit;
    justify-content: center;
}

.history-item-type {
    color: var(--front-3);
    grid-area: type;
    font-family: 'Figtree';
    cursor: inherit;
}

.history-item-delete-button {
    grid-area: dlt;
    height: calc(15px + 50% + 0.4vw);
    padding: 2px;
    aspect-ratio: 1/1;
    border: 1px solid transparent;
    border-radius: 30%;
}

.history-item-delete-button path {
    stroke: var(--front);
}

.history-item-delete-button:hover {
    background-color: var(--back-3);
}

.history-item-delete-button:hover path {
    stroke: var(--front-2);
}

.history-item-input-output {
    color: var(--front);
    font-family: 'Bahnschrift';
    cursor: inherit;
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
.dialog-buttons
{
    display: inline-flex;
    width: 100%;
    justify-content: right;
    gap: calc(4px + 0.5vw);
}
</style>
