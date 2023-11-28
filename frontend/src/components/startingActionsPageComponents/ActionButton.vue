<template>
  <div class="action-btn" :id="btnId" @click="navigateToLink">
    <div class="action-name">{{ name }}</div>
    <slot></slot>
    <label class="action-info">{{ info }}</label>
  </div>
</template>

  
<script>
export default {
  props: ['link', 'name', 'info', 'btnId'],
  methods: {
    navigateToLink() {
      if (this.link === "/projectCreation") {
        this.$emit('createNewProject');
      }
    }
  }
}
</script>
<style>
.action-btn {
  height: 100%;
  background-color: var(--back-2);
  border-radius: calc(0.15vh + 0.15vw + 2px);
  display: grid;
  display: grid;
  grid-template-rows: 1.8fr 1fr 1fr;
  grid-template-columns: 1fr 5fr;
  grid-template-areas:
    "icn hdr"
    "icn inf"
    "icn inf";
  cursor: pointer;
  max-height: 20vh;
  cursor: pointer;
}

.action-btn:hover {
  background-color: var(--back-3);
}

.action-name {
  grid-area: hdr;
  height: 100%;
  display: flex;
  align-items: end;
  color: var(--front-3);
  font-family: 'Bahnschrift';
  font-size: calc(0.92vw + 0.4vh + 6px);
  cursor: pointer;
}

.action-info {
  grid-area: inf;
  padding-right: 5px;
  margin-top: 0;
  color: var(--front);
  font-family: 'Bahnschrift';
  font-size: calc(0.62vw + 0.52vh + 2px);
  cursor: pointer;
}

.action-icon {
  grid-area: icn;
  align-self: center;
  width: 75%;
  aspect-ratio: 1/1;
  margin-left: 14%;
  cursor: pointer;
}

.action-icon path {
  stroke: var(--front-3);
}
</style>