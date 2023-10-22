
<template>
  <div class="tab-item" @click="selectTab" :class="{ 'active-tab': isActive }">
    <label>
      {{ tab.name }}
    </label>
    <svg class="tab-cross-btn" viewBox="0 0 24 24">
      <path
        d="M6.99486 7.00636C6.60433 7.39689 6.60433 8.03005 6.99486 8.42058L10.58 12.0057L6.99486 15.5909C6.60433 15.9814 6.60433 16.6146 6.99486 17.0051C7.38538 17.3956 8.01855 17.3956 8.40907 17.0051L11.9942 13.4199L15.5794 17.0051C15.9699 17.3956 16.6031 17.3956 16.9936 17.0051C17.3841 16.6146 17.3841 15.9814 16.9936 15.5909L13.4084 12.0057L16.9936 8.42059C17.3841 8.03007 17.3841 7.3969 16.9936 7.00638C16.603 6.61585 15.9699 6.61585 15.5794 7.00638L11.9942 10.5915L8.40907 7.00636C8.01855 6.61584 7.38538 6.61584 6.99486 7.00636Z" />
    </svg>
  </div>
</template>
<script>
export default {
  props: ['tab', 'isActive'],
  methods: {
    selectTab() {
      this.$emit('select', this.tab.id);
    }
  }
}
</script>
  
<style scoped>
.tab-item {
  font-size: calc(0.22vh + 0.3vw + 10px);
  font-weight: 400;
  background-color: var(--back-2);
  border: calc(0.04vw + 1px) solid var(--back-3);
  float: left;
  height: calc(0.8vh + 22px);
  display: inline-flex;
  padding-right: 0;
  transition: background-color 0.08s;
}

.tab-item path {
  display: none;
}

.tab-item:hover {
  background-color: var(--back-3);
}

.tab-item:hover path {
  display: block;
}

.tab-item label {
  margin-top: 3px;
  text-decoration: none;
  align-self: center;
  color: var(--front);
  font-family: 'Bahnschrift';
  display: flex;
  align-items: center;
  padding-left: calc(0.3vw + 5px);
}

.tab-cross-btn {
  height: calc(17px + 0.7vh);
  aspect-ratio: 1/1;
  align-self: center;
  margin-top: 0.08vh;
  margin-left: 3px;
}

.tab-cross-btn:hover path {
  fill: var(--front-3);
}

.tab-cross-btn path {
  fill: var(--front);
}

.active-tab label {
  color: var(--front-3);
}
</style>
  