<script lang="ts" setup>
import SendIcon from "~icons/mdi/send";
import TreeMenu from "@/components/partials/TreeMenu.vue";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { GetCurrentOids } from "../../wailsjs/go/main/App";
import { oidstorage } from "../../wailsjs/go/models";
import { ref } from "vue";

const showMessage = ref(false);

// TODO : actually read in OIDs from golang

let oids = Array<oidstorage.Oid>();

function TestThing() {
  GetCurrentOids().then((val) => {
    oids = val;
  });
  console.log("test");
}

let tree = {
  label: "root",
  nodes: [
    {
      label: "item1",
      nodes: [
        {
          label: "item1.1",
        },
        {
          label: "item1.2",
          nodes: [
            {
              label: "item1.2.1",
            },
          ],
        },
      ],
    },
    {
      label: "item2",
    },
  ],
};

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
      @click="TestThing"
    >
      <SendIcon height="20" width="20" />
    </button>
    <div class="flex justify-start text-left">
      <TreeMenu
        class=""
        :label="tree.label"
        :nodes="tree.nodes"
        :depth="0"
      ></TreeMenu>
    </div>
  </div>
</template>
