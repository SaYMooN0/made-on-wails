
<template>
  <label class="color-label non-selectable">
    {{ label }}:
    <div class="color-operating-container">

      <label class="color-pair">
        <template v-if="isEditing">
          <input v-model="editableColor" class="color-input" />
        </template>
        <template v-else>
          {{ color }}
        </template>
        <span :style="`background-color: ${color}`" class="color-div"></span>
      </label>

      <div class="change-color-button">
        <template v-if="isEditing">
          <IconButton @click="saveNewColor">
            <svg viewBox="0 0 24 24" fill="none">
              <path
                d="M10.33 16.08C10.2315 16.0805 10.1338 16.0612 10.0428 16.0235C9.9518 15.9857 9.86925 15.9301 9.8 15.86L6.68 12.75C6.53955 12.6094 6.46066 12.4188 6.46066 12.22C6.46066 12.0213 6.53955 11.8306 6.68 11.69C6.74895 11.6192 6.83136 11.563 6.92239 11.5246C7.01342 11.4862 7.11121 11.4664 7.21 11.4664C7.3088 11.4664 7.40659 11.4862 7.49762 11.5246C7.58864 11.563 7.67106 11.6192 7.74 11.69L10.33 14.27L16.42 8.18C16.5636 8.04131 16.7554 7.96378 16.955 7.96378C17.1546 7.96378 17.3464 8.04131 17.49 8.18C17.6305 8.32063 17.7093 8.51125 17.7093 8.71C17.7093 8.90875 17.6305 9.09938 17.49 9.24L10.86 15.86C10.7908 15.9301 10.7082 15.9857 10.6172 16.0235C10.5262 16.0612 10.4285 16.0805 10.33 16.08Z"
                fill="#000000" />
              <path
                d="M17.68 21.81H6.36C5.26466 21.81 4.21418 21.3749 3.43965 20.6004C2.66513 19.8258 2.23 18.7753 2.23 17.68V6.36C2.23 5.26466 2.66513 4.21417 3.43965 3.43965C4.21418 2.66512 5.26466 2.23 6.36 2.23H17.68C18.7753 2.23 19.8258 2.66512 20.6004 3.43965C21.3749 4.21417 21.81 5.26466 21.81 6.36V17.68C21.81 18.7753 21.3749 19.8258 20.6004 20.6004C19.8258 21.3749 18.7753 21.81 17.68 21.81ZM6.36 3.73C5.66248 3.73 4.99353 4.00709 4.50031 4.50031C4.00709 4.99353 3.73 5.66248 3.73 6.36V17.68C3.73 18.3775 4.00709 19.0465 4.50031 19.5397C4.99353 20.0329 5.66248 20.31 6.36 20.31H17.68C18.3775 20.31 19.0465 20.0329 19.5397 19.5397C20.0329 19.0465 20.31 18.3775 20.31 17.68V6.36C20.31 5.66248 20.0329 4.99353 19.5397 4.50031C19.0465 4.00709 18.3775 3.73 17.68 3.73H6.36Z"
                fill="#000000" />
            </svg>
          </IconButton>
          <IconButton @click="cancelEditing">
            <svg viewBox="0 0 24 24" fill="none">
              <path
                d="M17.66 21.81H6.34C5.24466 21.81 4.19417 21.3749 3.41965 20.6004C2.64512 19.8258 2.21 18.7753 2.21 17.68V6.36C2.21 5.26466 2.64512 4.21418 3.41965 3.43965C4.19417 2.66513 5.24466 2.23 6.34 2.23H17.66C18.7553 2.23 19.8058 2.66513 20.5803 3.43965C21.3549 4.21418 21.79 5.26466 21.79 6.36V17.68C21.79 18.7753 21.3549 19.8258 20.5803 20.6004C19.8058 21.3749 18.7553 21.81 17.66 21.81ZM6.34 3.73C5.64248 3.73 4.97353 4.00709 4.48031 4.50031C3.98709 4.99353 3.71 5.66248 3.71 6.36V17.68C3.71 18.3775 3.98709 19.0465 4.48031 19.5397C4.97353 20.0329 5.64248 20.31 6.34 20.31H17.66C18.3575 20.31 19.0265 20.0329 19.5197 19.5397C20.0129 19.0465 20.29 18.3775 20.29 17.68V6.36C20.29 5.66248 20.0129 4.99353 19.5197 4.50031C19.0265 4.00709 18.3575 3.73 17.66 3.73H6.34Z"
                fill="#000000" />
              <path
                d="M15.76 16.64C15.6614 16.641 15.5636 16.6221 15.4725 16.5843C15.3814 16.5465 15.2989 16.4906 15.23 16.42L7.58 8.77001C7.50631 8.70134 7.44721 8.61854 7.40622 8.52654C7.36523 8.43455 7.34318 8.33523 7.34141 8.23453C7.33963 8.13383 7.35815 8.0338 7.39588 7.94041C7.4336 7.84702 7.48974 7.76219 7.56096 7.69097C7.63218 7.61975 7.71701 7.56361 7.8104 7.52588C7.90379 7.48816 8.00382 7.46964 8.10452 7.47141C8.20522 7.47319 8.30454 7.49523 8.39654 7.53623C8.48854 7.57722 8.57134 7.63632 8.64 7.71001L16.29 15.36C16.4304 15.5006 16.5093 15.6913 16.5093 15.89C16.5093 16.0888 16.4304 16.2794 16.29 16.42C16.2208 16.4901 16.1382 16.5457 16.0472 16.5835C15.9562 16.6213 15.8585 16.6405 15.76 16.64Z"
                fill="#000000" />
              <path
                d="M8.19 16.56C8.09121 16.5622 7.99304 16.5438 7.90178 16.5059C7.81052 16.468 7.72817 16.4115 7.66 16.34C7.58924 16.2711 7.53301 16.1886 7.49461 16.0976C7.45621 16.0066 7.43643 15.9088 7.43643 15.81C7.43643 15.7112 7.45621 15.6134 7.49461 15.5224C7.53301 15.4314 7.58924 15.3489 7.66 15.28L15.31 7.63C15.4506 7.48955 15.6412 7.41066 15.84 7.41066C16.0387 7.41066 16.2294 7.48955 16.37 7.63C16.5087 7.77357 16.5862 7.96538 16.5862 8.165C16.5862 8.36462 16.5087 8.55643 16.37 8.7L8.72 16.34C8.65075 16.4101 8.5682 16.4657 8.47719 16.5035C8.38618 16.5412 8.28854 16.5605 8.19 16.56Z"
                fill="#000000" />
            </svg>
          </IconButton>
        </template>
        <template v-else>
          <IconButton @click="toggleEditing">
            <svg viewBox="0 0 24 24" fill="none">
              <path
                d="M11 4H7.2C6.0799 4 5.51984 4 5.09202 4.21799C4.71569 4.40974 4.40973 4.7157 4.21799 5.09202C4 5.51985 4 6.0799 4 7.2V16.8C4 17.9201 4 18.4802 4.21799 18.908C4.40973 19.2843 4.71569 19.5903 5.09202 19.782C5.51984 20 6.0799 20 7.2 20H16.8C17.9201 20 18.4802 20 18.908 19.782C19.2843 19.5903 19.5903 19.2843 19.782 18.908C20 18.4802 20 17.9201 20 16.8V12.5M15.5 5.5L18.3284 8.32843M10.7627 10.2373L17.411 3.58902C18.192 2.80797 19.4584 2.80797 20.2394 3.58902C21.0205 4.37007 21.0205 5.6364 20.2394 6.41745L13.3774 13.2794C12.6158 14.0411 12.235 14.4219 11.8012 14.7247C11.4162 14.9936 11.0009 15.2162 10.564 15.3882C10.0717 15.582 9.54378 15.6885 8.48793 15.9016L8 16L8.04745 15.6678C8.21536 14.4925 8.29932 13.9048 8.49029 13.3561C8.65975 12.8692 8.89125 12.4063 9.17906 11.9786C9.50341 11.4966 9.92319 11.0768 10.7627 10.2373Z"
                stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
          </IconButton>
        </template>

      </div>
    </div>
  </label>
</template>

<script>
import IconButton from './IconButton.vue';

export default {
  props: {
    label: String,
    color: String
  },
  data() {
    return {
      isEditing: false,
      editableColor: this.color
    };
  },
  methods: {
    toggleEditing() {
      this.cancelColorsEditing();
      this.isEditing = !this.isEditing;
    },
    saveNewColor() {
      if (this.isValidColor(this.editableColor)) {
        this.updateTheme(this.label, this.editableColor);
        this.isEditing = false;
      } else {
        this.showError(`Invalid color ${this.editableColor} for ${this.label}`);
      }
    },
    cancelEditing() {
      this.editableColor = this.color;
      this.isEditing = false;
    },
    isValidColor(color) {
      const regex = /^#[0-9A-Fa-f]{6}$/;
      return regex.test(color);
    }
  },
  components: {
    IconButton
  },
  inject: ['cancelColorsEditing', 'updateTheme', 'showError'],
};
</script>

<style scoped>
.color-label {
  color: black;
  font-size: calc(0.35vh + 0.5vw + 10px);
  font-weight: 600;
  font-family: 'Figtree';
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0;
}

.color-div {
  height: calc(0.5vh + 0.4vw + 10px);
  aspect-ratio: 1/1;
  border-radius: 10%;
}

.color-operating-container {
  width: calc(40% + 54px);
  display: grid;
  grid-template-columns: 1fr calc(0.7vh + 0.6vw + 38px);
  gap: calc(0.6vw + 5px);

}

.color-pair {
  display: grid;
  grid-template-columns: 1fr calc(0.5vh + 0.4vw + 10px);
  gap: calc(0.4vw + 1px);
  align-items: center;
  font-size: calc(0.35vh + 0.5vw + 8px);
  text-align: right;

}

.color-input {
  width: calc(1.8vh + 2vw + 34px);
  font-size: calc(0.35vh + 0.5vw + 8px);
  background-color: #e2e6ea;
  font-weight: 600;
  font-family: 'Figtree';
  text-align: center;
  outline: none;
  padding: 2px;
  border: 1px solid transparent;
  border-radius: 2px;
  justify-self: end;
}

.color-input:focus {
  background-color: #d8dee2;
}

.change-color-button {
  height: 100%;
  aspect-ratio: 2.2/1;
  display: flex;
  justify-content: center;
}
</style>
  