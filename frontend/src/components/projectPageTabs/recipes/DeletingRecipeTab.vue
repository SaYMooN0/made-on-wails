<template>
     <div class="tab-container">
            <DefCheckBox labelText="From"></DefCheckBox>
            <DefCheckBox labelText="To"></DefCheckBox>

            item:<input class="default-input" type="text">
            type:<input class="default-input" type="text">
        </div>
</template>
<script>
import DefCheckBox from '../../default/DefCheckBox.vue';
export default {
    components:
    {
        DefCheckBox
    },
}
</script>
<style scoped>
.tab-container
{

}
</style>