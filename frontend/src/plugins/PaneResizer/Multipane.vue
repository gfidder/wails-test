<script lang="ts" setup>
import { computed, ref } from "vue";

const LAYOUT_HORIZONTAL = "horizontal";
const LAYOUT_VERTICAL = "vertical";

const isResizing = ref(false);

const props = defineProps({
  layout: { type: String, default: LAYOUT_HORIZONTAL },
});

const classnames = computed(() => {
  return [
    "multipane",
    "layout-" + props.layout.slice(0, 1),
    isResizing.value ? "is-resizing" : "",
  ];
});

const cursor = computed(() => {
  return isResizing.value
    ? props.layout == LAYOUT_VERTICAL
      ? "col-resize"
      : "row-resize"
    : "";
});

const userSelect = computed(() => {
  return isResizing.value ? "none" : "";
});

// function onMouseDown({
//   target: resizer,
//   pageX: initialPageX,
//   pageY: initialPageY,
// }): void {
//   if (resizer.className && resizer.className.match("multipane-resizer")) {
//     let self = this;
//     let { $el: container, layout } = self;

//     let pane = resizer.previousElementSibling;
//     let { offsetWidth: initialPaneWidth, offsetHeight: initialPaneHeight } =
//       pane;

//     let usePercentage = !!(pane.style.width + "").match("%");

//     const { addEventListener, removeEventListener } = window;

//     const resize = (initialSize, offset = 0) => {
//       if (layout == LAYOUT_VERTICAL) {
//         let containerWidth = container.clientWidth;
//         let paneWidth = initialSize + offset;

//         return (pane.style.width = usePercentage
//           ? (paneWidth / containerWidth) * 100 + "%"
//           : paneWidth + "px");
//       }

//       if (layout == LAYOUT_HORIZONTAL) {
//         let containerHeight = container.clientHeight;
//         let paneHeight = initialSize + offset;

//         return (pane.style.height = usePercentage
//           ? (paneHeight / containerHeight) * 100 + "%"
//           : paneHeight + "px");
//       }
//     };

//     isResizing.value = true;

//     let size = resize();

//     self.$emit("paneResizeStart", pane, resizer, size);

//     const onMouseMove = function ({ pageX, pageY }) {
//       size =
//         layout == LAYOUT_VERTICAL
//           ? resize(initialPaneWidth, pageX - initialPageX)
//           : resize(initialPaneHeight, pageY - initialPageY);

//       self.$emit("paneResize", pane, resizer, size);
//     };

//     const onMouseUp = function () {
//       size =
//         layout == LAYOUT_VERTICAL
//           ? resize(pane.clientWidth)
//           : resize(pane.clientHeight);

//       isResizing.value = false;

//       removeEventListener("mousemove", onMouseMove);
//       removeEventListener("mouseup", onMouseUp);

//       self.$emit("paneResizeStop", pane, resizer, size);
//     };

//     addEventListener("mousemove", onMouseMove);
//     addEventListener("mouseup", onMouseUp);
//   }
// }
</script>

<template>
  <div :class="classnames" @mousedown="">
    <slot></slot>
  </div>
</template>

<style lang="scss">
.multipane {
  display: flex;

  &.layout-h {
    flex-direction: column;
  }

  &.layout-v {
    flex-direction: row;
  }
}

.multipane > div {
  position: relative;
  z-index: 1;
}

.multipane-resizer {
  display: block;
  position: relative;
  z-index: 2;
}

.layout-h > .multipane-resizer {
  width: 100%;
  height: 10px;
  margin-top: -10px;
  top: 5px;
  cursor: row-resize;
}

.layout-v > .multipane-resizer {
  width: 10px;
  height: 100%;
  margin-left: -10px;
  left: 5px;
  cursor: col-resize;
}
</style>
