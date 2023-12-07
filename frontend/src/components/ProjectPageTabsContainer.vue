<template>
  <div class="tabs-container">
    <VueDraggableNext :list="tabs" @update:list="$emit('update:tabs', $event)" class="tab-list">
      <ProjectPageTab v-for="tab in tabs" :key="tab.id" :tab="tab" @select="$emit('update:activeTabId', tab.id)" 
      @remove="$emit('remove', tab.id)" :isActive="activeTabId === tab.id" />
    </VueDraggableNext>
  </div>
</template>

<script>
import { VueDraggableNext } from 'vue-draggable-next';
import ProjectPageTab from './ProjectPageTab.vue';

export default {
  components: {
    VueDraggableNext,
    ProjectPageTab
  },
  props: ['tabs', 'activeTabId']
}
</script>

<style scoped>
.tabs-container {
  width: 100%;
  margin-left: 0;
  padding: 0;
  list-style-type: none;
  white-space: nowrap;
  overflow-y: auto;
  margin-top: calc(6px + 0.4vh);
  background-color: var(--back-2);
  height: calc(0.8vh + 25px);
  padding: 0;
}

.tabs-container::-webkit-scrollbar {
  height: calc(40% + 9px);
}

.tabs-container::-webkit-scrollbar-thumb {
  border-radius: 5px;
  background-color: var(--front-2);
}

.tabs-container::-webkit-scrollbar {
  width: calc(0.3vw + 6px);
  max-width: calc(9px);
}

.tabs-container::-webkit-scrollbar-thumb:hover {
  background-color: var(--bright);
}
</style>