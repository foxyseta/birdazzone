<script setup lang=ts>
import { ref, onMounted } from 'vue'
import { createPopper } from "@popperjs/core";
import ApiRepository from '../api/api-repository';
import { SemipolarSpinner } from 'epic-spinners';

const error = ref<boolean>(false)
const loading = ref<boolean>(true)
const props = defineProps<{ id: string}>()
const nFail = ref<number>(0)
const nSucc = ref<number>(0)
const fail = ref<number>(0)
const success = ref<number>(0)
const nAttempts = ref<number>(0)
const canvas = ref(null)
const see = ref(false)
const popoverRef = ref(null)

const fetchData = async () => {
  const response = await ApiRepository.getResults(props.id)
  
  if (response.esit) {
    nFail.value = response.data!.negatives;             // failed attempts
    nSucc.value = response.data!.positives;             // succeded attempts
    nAttempts.value = nFail.value + nSucc.value;        // total attempts

    success.value = (() => {    // percentage 
      if (nAttempts.value == 0)
        return 0;
      else
        return nSucc.value / nAttempts.value;   // x : 1 = nSucc : nAttempts
    })();

    fail.value = 1-success.value;

  } else {
    error.value = true
  }
}

onMounted(async () => {
  loading.value = true

  await fetchData()

  loading.value = false

  if (!canvas.value)
    return

  // @ts-ignore
  const context = canvas.value.getContext('2d');
  // @ts-ignore
  const centerX = canvas.value.width / 2;
  // @ts-ignore
  const centerY = canvas.value.height / 2;
  const radius = 50;

  const spaceBetween = 0.33;
  const start = 3/2*Math.PI;

  const succPerc = (2 * Math.PI - spaceBetween) *success.value;     // 2PI : 1 =  x : success
  const failPerc = (2 * Math.PI - spaceBetween) - succPerc;

  context.beginPath();    // success
  context.lineCap = 'round';
  context.lineWidth = 13;
  context.arc(centerX, centerY, radius, start, start + succPerc, false);
  context.strokeStyle = '#1eb980';
  context.stroke();
  context.closePath();

  context.beginPath();    // fail
  context.arc(centerX, centerY, radius, start + succPerc + spaceBetween, start + succPerc + failPerc, false);
  context.strokeStyle = '#ff937f';
  context.lineWidth = 13;
  context.stroke();
  context.closePath();
});


const popover = () => {
  see.value = true;
  createPopper(canvas.value!, popoverRef.value!, {
    placement: 'top-end'
  });
  setTimeout(() => see.value = false, 2000);
}
</script>

<template>

  <div class="bg-foreground font-semibold text-lg rounded-lg p-4 place-self-center z-0" >

    <div v-if="loading">
      <semipolar-spinner :animation-duration="2000" :size="35" color="#1eb980" />
    </div>
    <div v-show="!loading && !error" class="flex items-center">
      <div class="items-center mx-3">
        <div class="text-white m-2">{{nAttempts}} tried</div>
        <div class="m-2"><div class="text-lgreen inline">{{nSucc}}</div><div class="text-white inline"> succeded</div></div>
        <div class="m-2"><div class="text-lred inline">{{nFail}}</div><div class="text-white inline"> failed</div></div>
      </div>

      <div class="relative z-10">
        <canvas @mouseover="popover()" class="mr-3 z-10" ref="canvas" width="150" height="150"></canvas>
        <div ref="popoverRef" v-bind:class="{'hidden': !see, 'block': see}"
          class="absolute inset-0 text-sm font-semibold text-center rounded-lg bg-lblack p-2">
          <div class="items-center">
              <div class="text-lred">failed</div>
              <div class="text-lgreen">succeded</div>
          </div>
        </div>
        <div class="bottom-0 left-0 w-full ml-6">
          <div class="space-x-6">
            <div class="text-lred inline">{{ (fail * 100).toFixed(0) }}%</div>
            <div class="text-lgreen inline">{{ (success * 100).toFixed(0) }}%</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
