<template>
  <h1>Project page</h1>
  <h1>{{JSON.stringify(project)}}</h1>
  <h1>End of project page</h1>

  <button @click="backToMain">Назад</button>
</template>
  
<script>
import { CurrentProject } from "../../../wailsjs/go/projectrelated/ProjectManager";
export default {
  data() {
    return {
      project: {}
    };
  },
  created() {
    this.fetchProject();
  },
  methods: {
    backToMain() {
      this.$emit('goBack');
    },
    fetchProject() {
      CurrentProject().then((projectToFetch) => {
                this.project = projectToFetch;
                console.log(projectToFetch);
                console.log(this.project);
            });
    },
  },



}
</script>