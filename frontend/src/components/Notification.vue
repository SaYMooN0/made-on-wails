<template>
    <div :class="['notification', notificationClass]">
      {{ message }}
    </div>
  </template>
  
<script>
export default {
  data() {
    return {
      visible: false,
      message: '',
    };
  },
  computed: {
    notificationClass() {
      return {
        'notification-visible': this.visible,
        'notification-hidden': !this.visible
      };
    }
  },
  methods: {
    show(message) {
      this.message = message;
      this.visible = true;
      setTimeout(() => {
        this.visible = false;
      }, 3000);
    },
  },
};
</script>



  
<style>
.notification {
    position: fixed;
    top: calc(10px + 2vh);
    left: 50%;
    transform: translateX(-50%) translateY(-50px);
    background-color: #ffcc00;
    color: #333;
    padding: 10px 20px;
    border-radius: 5px;
    box-shadow: 0 2px 15px rgba(0, 0, 0, 0.1);
    opacity: 0;
    transition: all 0.5s ease;
    z-index: 1000;
}

.notification-visible {
    opacity: 1;
    transform: translateX(-50%) translateY(0);
}

.notification-hidden {
    opacity: 0;
    transform: translateX(-50%) translateY(-20px);
}</style>
  