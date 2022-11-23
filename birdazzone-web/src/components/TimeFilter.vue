<script lang=ts>
import { ref } from 'vue';
import { defineComponent } from 'vue';
import DatePicker from 'vue-datepicker-next';
import 'vue-datepicker-next/index.css';

export default defineComponent({
  components: { DatePicker },
  data() {
    return {
      startDate: null,
      startTime: null,
      endDate: null,
      endTime: null,

      openSD: false,
      openST: false,
      openED: false,
      openET: false,
    };
  },
  methods: {
    closeSD() {
      this.openSD = false;
    },
    closeST() {
      this.openST = false;
    },
    closeED() {
      this.openED = false;
    },
    closeET() {
      this.openET = false;
    },
  },
});

</script>

<template>
  <div class="bg-foreground font-semibold text-lg rounded-lg m-4 p-4 place-self-center">

    <!--<div v-if="loading">
      <semipolar-spinner :animation-duration="2000" :size="35" color="#1eb980" />
    </div>-->
    <div class="flex flex-col">
      <label for="startDateTime" class="justify-self-start text-white">Start</label>
      <div class="flex flex-col justify-self-start" id="startDateTime">
        <date-picker
          class = "m-2"
          v-model:value = "startDate"
          v-model:open="openSD"
          :clearable = "true"
          type = "date"
          placeholder="select start date">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeSD()">close</button>
            </template>
        </date-picker>
        <date-picker
          class = "m-2"
          v-model:value = "startTime" 
          v-model:open="openST"
          :clearable = "true"
          type = "time"
          placeholder="select start time">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeST()">close</button>
            </template>
        </date-picker>
      </div>
    </div>

    <div class="flex flex-col">
      <label for="startDateTime" class="justify-self-start text-white">End</label>
      <div class="flex flex-col justify-self-start" id="startDateTime">
        <date-picker
          class = "m-2"
          v-model:value = "endDate" 
          v-model:open="openED"
          :clearable = "true"
          type = "date"
          placeholder="select end date">
          <template #footer>
            <button class="mx-btn mx-btn-text" @click="closeED()">close</button>
          </template>
        </date-picker>
        <date-picker
          class = "m-2"
          v-model:value = "endTime" 
          v-model:open="openET"
          :clearable = "true"
          type = "time"
          placeholder="select end time">
            <template #footer>
              <button class="mx-btn mx-btn-text" @click="closeET()">close</button>
            </template>
        </date-picker>
      </div>
    </div>

  </div>
</template>
