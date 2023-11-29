<template>
    <div class="page-container">


        <h1 class="new-theme-name">
            New theme name:
            <input type="text" />
        </h1>
        <div class="color-inputs-blocks-container">
            <div class="color-inputs-block">
                <label class="color-group-label">Back Colors:</label>
                <NewThemeColorInput />
                <NewThemeColorInput />
                <NewThemeColorInput />
            </div>
            <div class="color-inputs-block">
                <label class="color-group-label">Front Colors:</label>
                <NewThemeColorInput />
                <NewThemeColorInput />
                <NewThemeColorInput />
            </div>
            <div class="color-inputs-block">
                <label class="color-group-label">Bright Colors:</label>
                <NewThemeColorInput />
                <NewThemeColorInput />
                <NewThemeColorInput />
            </div>
            <div class="color-inputs-block">
                <label class="color-group-label">Warning Colors:</label>
                <NewThemeColorInput />
                <NewThemeColorInput />
            </div>
        </div>
        <GoBackButton @goBack="backToMain" />
    </div>
</template>
<script>
import NewThemeColorInput from '../newThemePageComponents/NewThemeColorInput.vue';
import GoBackButton from '../GoBackButton.vue';
import { GetHexDefaultDark } from '../../../wailsjs/go/themerelated/ThemeCollection'
export default {
    data() {
        return {
            newTheme: {},
        }

    },
    computed: {
        themeValues() {
            return Object.fromEntries(
                Object.entries(this.newTheme).filter(([key]) => key.endsWith("Color"))
            );
        },
    },
    components:
    {
        NewThemeColorInput,
        GoBackButton
    },
    methods: {
        backToMain() {
            this.$emit('goBack');
        }
    },
    created() {
        GetHexDefaultDark().then((theme) => {
            this.newTheme = theme;
        });
    }
};
</script>
<style>
.page-container {
    display: grid;
    grid-template-rows: auto 1fr;
    height: 100vh;
    padding-left: calc(1vh + 0.4vw + 11px);
    padding-right: calc(1vh + 0.4vw + 11px);
}

.new-theme-name {
    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(1.4vh + 1.4vw + 11px);
}

.theme-name-input {}

.color-inputs-blocks-container {
    background-color: aqua;
    height: 90%;
    width: 100%;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: calc(0.4vh + 1vw);
    align-items: center;
}

.color-inputs-block {
    height: 82%;
    width: 100%;
    background-color: #f5f5f5;
    border-radius: 2%;
    display: grid;
    grid-template-rows: auto 1fr 1fr 1fr;
}
</style>