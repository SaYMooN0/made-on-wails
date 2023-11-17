<template>
  <div class="tabs-container">
    <button @click="setActiveTab('visual')" class="tab-button">Visual</button>
    <button @click="setActiveTab('hotkeys')" class="tab-button">Hotkeys</button>
    <button @click="setActiveTab('others')" class="tab-button">Others</button>
  </div>
  <div class="tabs-content">
    <VisualSettingsTab v-if="activeTab === 'visual'" />
    <HotKeysSettingsTab v-if="activeTab === 'hotkeys'" />
    <OtherSettingsTab v-if="activeTab === 'others'" />
  </div>
  <button @click="backToMain">Назад</button>
</template>

<script>
import HotKeysSettingsTab from './../settingsPageComponents/Tabs/HotKeysSettingsTab.vue';
import OtherSettingsTab from './../settingsPageComponents/Tabs/OtherSettingsTab.vue';
import VisualSettingsTab from './../settingsPageComponents/Tabs/VisualSettingsTab.vue';

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
    }
  },
  components: {
    HotKeysSettingsTab,
    OtherSettingsTab,
    VisualSettingsTab
  }
}
</script>
<style scoped>
.tabs-container {
  position: absolute;
  width: 100vw;
  display: grid;
  grid-template-columns: auto auto auto 1fr;
  color: darkblue;
}
.tab-button{
  font-size: calc(0.22vh + 0.2vw + 12px);
  font-weight: 400;
  background-color: var(--back-2);
  border: calc(0.04vw + 1px) solid transparent;
  float: left;
  height: calc(0.8vh + 22px);
  align-items: center;
  display: inline-flex;
  color: var(--front);
}
.tab-button:hover{
  background-color: var(--back-3);
  color: var(--front-2);
}
</style>