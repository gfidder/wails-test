<script lang="ts" setup>
import SideBarIcon from "./SideBarIcon.vue";
import Divider from "./partials/Divider.vue";
import YouTubeIcon from "~icons/mdi/youtube";
import BlenderSoftwareIcon from "~icons/mdi/blender-software";
import TreeMenu from "./partials/TreeMenu.vue";

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
    <SideBarIcon>
      <YouTubeIcon height="28" width="28" />
    </SideBarIcon>
    <Divider />
    <SideBarIcon>
      <BlenderSoftwareIcon height="29" width="28" />
    </SideBarIcon>
    <Divider />
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
