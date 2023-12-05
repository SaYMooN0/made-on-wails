<template>
    <div :class="['notification', notificationClass, { error: isError }, 'non-selectable']"
         @mouseenter="onMouseEnter"
         @mouseleave="onMouseLeave">
      {{ message }}
      <span class="close-button" @click="closeNotification">&times;</span>
    </div>
  </template>
  
<script>
export default {
    data() {
        return {
            visible: false,
            isError: false,
            message: '',
            hideTimeout: null,
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
        show(message, isError = false) {
            this.message = message;
            this.visible = true;
            this.isError = isError;
            this.startHideTimer();
        },
        closeNotification() {
            this.visible = false;
            clearTimeout(this.hideTimeout); 
        },
        startHideTimer() {
            this.hideTimeout = setTimeout(() => {
                this.visible = false;
            }, 3000);
        },
        setTimeoutToOneSec()
        {
            this.hideTimeout = setTimeout(() => {
                this.visible = false;
            }, 1000);
        },
        onMouseEnter() {
            clearTimeout(this.hideTimeout);
        },
        onMouseLeave() {
            this.setTimeoutToOneSec();
        }
    },

};
</script>



  
<style>
.notification {
    position: fixed;
    top: calc(8px + 5vh);
    left: 50%;
    background-color: var(--back);
    color: var(--front);
    font-family: Bahnschrift;
    font-weight: 700;
    font-size: calc(9px + 0.7vw + 0.6vh);
    padding: 10px 20px;
    border-radius: calc(3px + 0.15vh + 0.15vw);
    box-shadow: 1px 4px 15px rgba(0.2, 0.2, 0.2, 0.35);
    opacity: 0;
    transition: all 0.5s ease;
    z-index: 1000;
    width: calc(150px + 50%);
}

.notification-visible {
    opacity: 1;
    transform: translateX(-50%) translateY(0);
}

.notification-hidden {
    opacity: 0;
    transform: translateX(-50%) translateY(-20px);
}

.error {
    color: var(--warning-main) !important;
}

.close-button {
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translateY(-50%);
    cursor: pointer;
    color: var(--front);
    font-size: calc(14px + 0.9vw + 0.8vh);
}

.close-button:hover {
    color: var(--bright);
}
</style>
  