import { createApp } from "vue";
import App from "./App.vue";
import router from "./routers";

// import "vfonts/Lato.css";
// import "vfonts/FiraCode.css";

createApp(App).use(router).mount("#app");
