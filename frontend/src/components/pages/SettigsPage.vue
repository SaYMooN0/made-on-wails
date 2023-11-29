<template>
  <div class="settings-container">
    <div class="tabs-container">
      <button @click="setActiveTab('visual')" class="tab-button">Visual</button>
      <button @click="setActiveTab('hotkeys')" class="tab-button">Hotkeys</button>
      <button @click="setActiveTab('others')" class="tab-button">Others</button>
    </div>
    <div class="tabs-content">
      <VisualSettingsTab v-if="activeTab === 'visual'" @createNewTheme="createNewTheme" />
      <HotKeysSettingsTab v-if="activeTab === 'hotkeys'" />
      <OtherSettingsTab v-if="activeTab === 'others'" />
    </div>
  </div>
  <GoBackButton @goBack="backToMain"/>
</template>

<script>
import HotKeysSettingsTab from './../settingsPageComponents/Tabs/HotKeysSettingsTab.vue';
import OtherSettingsTab from './../settingsPageComponents/Tabs/OtherSettingsTab.vue';
import VisualSettingsTab from './../settingsPageComponents/Tabs/VisualSettingsTab.vue';
import GoBackButton from '../GoBackButton.vue';
export default {
  data() {
    return {
      activeTab: 'visual'
    }
  },
  methods: {
    backToMain() {
      this.$emit('goBack');
    },
    setActiveTab(tabName) {
      this.activeTab = tabName;
    },
    createNewTheme() {
      this.$emit('goToNewTheme');
    }
  },
  components: {
    HotKeysSettingsTab,
    OtherSettingsTab,
    VisualSettingsTab,
    GoBackButton
  }
}
</script>
<style scoped>
.settings-container {
  display: grid;
  grid-template-rows: calc(1vh + 32px + 0.1vw) 1fr;
  height: 100vh;
}

.tabs-container {
  height: calc(100% - 2px);
  width: calc(100% - 10px);
  padding-left: 10px;
  display: grid;
  grid-template-columns: auto auto auto 1fr;
  align-items: center;
  border-bottom: 2px var(--back-2) solid;
  ;
}

.tabs-content {
  width: 100%;
  height: 100%;
}

.tab-button {
  font-size: calc(0.22vh + 0.2vw + 12px);
  font-weight: 600;
  background-color: transparent;
  border: 1px solid transparent;
  border-radius: calc(3px + 0.09vw + 0.1vh);
  align-items: center;
  color: var(--front);
  height: 80%;
  padding: 0 calc(4px + 0.12vw + 0.7vh) 0 calc(4px + 0.12vw + 0.7vh);
  font-family: 'Figtree';
}

.tab-button:hover {
  background-color: var(--back-3);
  color: var(--front-2);
}

.go-back-button {
  position: absolute;
  right: calc(8px + 1vw);
  bottom: calc(6px + 2%);
  height: calc(2vh + 22px + 0.2vw);
  width: calc(4vw + 60px + 2vh);
  background-color: var(--bright-2);
  font-family: 'Figtree';
  font-size: calc(0.55vh + 0.55vw + 10px);
  font-weight: 300;
  color: var(--front);
  border: 1px solid transparent;
  border-radius: calc(3px + 0.09vw + 0.1vh);
  cursor: pointer;
  transition: 0.08s;
}

.go-back-button:hover {
  background-color: var(--bright-3);
}
</style>