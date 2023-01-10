<script lang="ts" setup>
import SendIcon from "~icons/mdi/send";
import TreeMenu from "@/components/partials/TreeMenu.vue";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { GetCurrentOids } from "../../wailsjs/go/main/App";
import { OidTree, TreeSorter } from "../utils/treeBuilder";
import { ref, reactive, onBeforeMount } from "vue";

const showMessage = ref(false);
const updateCounter = ref(0);

onBeforeMount(ReloadMibTree);

async function ReloadMibTree() {
  otherOidTree.oidTree = new TreeSorter(await GetCurrentOids()).createOidTree();
  updateCounter.value++;
}

const otherOidTree = reactive({
  oidTree: { name: "oids loading...", oid: "place2" } as OidTree,
});

EventsOn("mibsLoaded", () => {
  showMessage.value = !showMessage.value;
});
</script>

<template>
  <div
    class="fixed top-0 left-0 w-1/3 h-screen flex flex-col bg-white dark:bg-gray-900 shadow-lg"
  >
    <div
      id="title"
      class="bg-gradient-to-r from-indigo-600 to-white h-14 text-left flex"
    >
      <h2 class="text-2xl text-black flex items-center">SNMP Oids</h2>
    </div>
    <h2 v-if="showMessage">Message</h2>
    <button
      class="relative flex items-center justify-center h-12 w-12 mt-2 mb-2 mx-auto bg-gray-400 hover:bg-green-600 dark:bg-gray-800 text-green-500 hover:text-white hover:rounded-xl rounded-3xl transition-all duration-300 ease-linear cursor-pointer shadow-lg group"
      @click="ReloadMibTree"
    >
      <SendIcon height="20" width="20" />
    </button>
    <div class="flex justify-start text-left">
      <TreeMenu
        :key="updateCounter"
        class=""
        :node="otherOidTree.oidTree"
        :depth="0"
      ></TreeMenu>
    </div>
  </div>
</template>
