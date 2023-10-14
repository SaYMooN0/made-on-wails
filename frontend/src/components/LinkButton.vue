<template>
  <div class="link-button" :class="iconColor">
    <div class="icon-container">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: ['iconColor'],
}
</script>
<style scoped>
.link-button {
  background-color: var(--back-2);
  border: 0 solid transparent;
  height: calc(24px + 0.7vw + 0.7vh);
  aspect-ratio: 1/1;
  cursor: pointer;
  border-radius: calc(0.17vw + 0.17vh + 4px);
  display: flex;
  align-items: center;
  justify-content: center;

  transition: 0.04s;

}

.link-button:hover {
  background-color: var(--back-3);

}

.fill:hover {
  fill: var(--front-3);

}

.stroke:hover {
  stroke: var(--front-3);

}

.stroke {
  stroke: var(--front);
}

.fill {
  fill: var(--front);
}

.icon-container {
  height: 92%;
  aspect-ratio: 1/1;
  overflow: hidden;
  transition: 0.08s;

}

.icon-container>svg {
  max-width: 100%;
  max-height: 100%;
}</style>