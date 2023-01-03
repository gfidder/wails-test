<script lang="ts" setup>
import SideBarIcon from "./SideBarIcon.vue";
import Divider from "./partials/Divider.vue";
import YouTubeIcon from "~icons/mdi/youtube";
import BlenderSoftwareIcon from "~icons/mdi/blender-software";
</script>

<template>
  <div
    class="fixed top-0 left-0 h-screen flex flex-col bg-white dark:bg-gray-900 shadow-lg"
  >
    <SideBarIcon>
      <YouTubeIcon height="28" width="28" />
    </SideBarIcon>
    <Divider />
    <SideBarIcon>
      <BlenderSoftwareIcon height="29" width="28" />
    </SideBarIcon>
  </div>
</template>
