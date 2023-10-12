<template>
  <div class="project-item" @click="projectItemClicked(project.PathToFolder)">
    <label class="project-text-name">{{ project.Name }}</label>
    <label class="project-text-path">{{ project.FullPath }}</label>
    <svg class="project-more-button" fill="#000000" version="1.1" viewBox="0 0 342.382 342.382" xml:space="preserve">
                <path d="M45.225,125.972C20.284,125.972,0,146.256,0,171.191c0,24.94,20.284,45.219,45.225,45.219c24.926,0,45.219-20.278,45.219-45.219C90.444,146.256,70.151,125.972,45.225,125.972z" />
                <path d="M173.409,125.972c-24.938,0-45.225,20.284-45.225,45.219c0,24.94,20.287,45.219,45.225,45.219c24.936,0,45.226-20.278,45.226-45.219C218.635,146.256,198.345,125.972,173.409,125.972z" />
                <path d="M297.165,125.972c-24.932,0-45.222,20.284-45.222,45.219c0,24.94,20.29,45.219,45.222,45.219 c24.926,0,45.217-20.278,45.217-45.219C342.382,146.256,322.091,125.972,297.165,125.972z" />
            </svg>
    <label class="project-text-version">{{ project.Loader }} {{ project.Version }}</label>
  </div>
</template>
  
<script>
export default {
  props: ['project'],
  methods: {
    projectItemClicked(path) {
    }
  }
}
</script>
<style scoped>
.project-item {
  min-height: 60px;
  height: calc(20px + 4.8vh + 2.6vw);
  margin-left: 0.5%;
  width: 96%;
  margin-top: 0.6%;
  background-color: var(--back);
  display: grid;
  grid-template-rows: 1fr 1fr;
  grid-template-columns: 5fr 8fr calc(6vw + 48px) calc(3.8vw + 12px);
  grid-template-areas:
    "name name vrsn more"
    "path path path path";
  border-radius: calc(0.12vw + 0.07vh + 3px);
  cursor: pointer;
}

.project-item:hover {
  background-color: var(--back-2);
}

.project-text-name {
  grid-area: name;
  height: 100%;
  color: var(--front-3);
  font-family: 'Bahnschrift';
  font-size: calc(1.35vh + 1.3vw + 4px);
  display: flex;
  align-items: center;
  font-weight: 700;
  margin-top: 0.4vh;
  margin-left: 1.8%;
  width: 98%;
  cursor: pointer;
}

.project-text-path {
  grid-area: path;
  height: 80%;
  color: var(--front);
  font-family: 'Figtree';
  font-size: calc(0.5vh + 0.5vw + 4px);
  text-align: left;
  margin-top: 2px;
  margin-left: 1.8%;
  width: 92%;
  word-wrap: break-word;
  cursor: pointer;
}

.project-text-version {
  grid-area: vrsn;
  color: var(--front-2);
  font-family: 'Bahnschrift';
  margin-top: 7px;
  font-size: calc(0.9vh + 0.9vw + 3px);
  text-align: right;
  font-weight: 600;
  cursor: pointer;
}

.project-more-button {
  margin-left: 24%;
  margin-top: calc(1px + 5%);
  grid-area: more;
  width: calc(1.3vw + 9px);
  aspect-ratio: 1/1;
  cursor: pointer;
  cursor: pointer;
}

.project-more-button path {
  fill: var(--front-2);
}

.project-more-button:hover path {
  fill: var(--bright-2);

}

.corrupted-sign {
  margin-top: 8%;
  grid-area: vrsn;
  width: calc(6px + 40%);
  aspect-ratio: 1/1;
  cursor: pointer;
  position: relative;
  right: -26%;
}

.corrupted-sign-line {
  stroke: var(--warning-main);
}

.corrupted-sign:hover .corrupted-sign-line {
  stroke: var(--warning-bright);
  stroke-width: 2.4;

}
</style>