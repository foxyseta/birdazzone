<script setup lang="ts">
import ApiRepository from '../api/api-repository';
import type { Coordinates } from '../api/interfaces/tweet';
import { onBeforeMount, ref } from 'vue'
import ErrorWidget from './ErrorWidget.vue'
import FilterMap from './FilterList.vue'
import { SemipolarSpinner } from 'epic-spinners';

const props = defineProps<{ gameId: string, key: number, from: string | null, to: string | null }>()

const ROME = [ 12.706374170037495, 42.21140846575139 ]

const loading = ref<boolean>(false)
const center = ref<number[]>(ROME)
const zoom = ref<number>(6)
const rotation = ref<number>(0)
const fillColor = ref<string>('#1eb980')
const strokeColor = ref<string>('white')
const strokeWidth = ref<number>(3)
const radius = ref<number>(10)
const projection = ref<string>("EPSG:4326")
const coordinates = ref<number[][]>([]) // [long, lat]
const from = ref<string | null>(props.from)
const to = ref<string | null>(props.to)

const unpackCoordinates = (coordinates: Coordinates) => [coordinates.longitude, coordinates.latitude]

const fetchCoordinates = async () => {
  if(!from.value || !to.value){
    const response = await ApiRepository.getListOfGuesser(props.gameId, "1", "100")
    if (response.esit) {
      // @ts-ignore
      coordinates.value = response.data!.entries.map(tweet => tweet.coordinates).filter(c => c).map(unpackCoordinates) // Remove undefined and nulls
    } else {
      if (response.statusCode === 204) {
      error.value = true
        errorTitle.value = 'There are no tweets!'
        errorText.value = 'No game instance has been played!'
      }
      else {
        error.value = true
        errorTitle.value = 'Error!'
        errorText.value = 'something went wrong!'
      }
    }
  }
  else{
    const response = await ApiRepository.getListOfGuesserFiltered(props.gameId, from.value.toString(), to.value.toString(), "1", "100")
    if (response.esit) {
      // @ts-ignore
      coordinates.value = response.data!.entries.map(tweet => tweet.coordinates).filter(c => c).map(unpackCoordinates) // Remove undefined and nulls
    } else {
      if (response.statusCode === 204) {
        error.value = true
        errorTitle.value = 'There are no tweets!'
        errorText.value = 'No game instance has been played!'
      }
      else {
        error.value = true
        errorTitle.value = 'Error!'
        errorText.value = 'something went wrong!'
      }
    }
  }
}

onBeforeMount(async () => {
  loading.value = true
  await fetchCoordinates()
  loading.value = false
})

</script>

<template>
  <!-- Error -->
  <div v-if="error" class="flex justify-center items-center w-full">
    <ErrorWidget :open="true" :title="errorTitle" :text="errorText" />
  </div>
  
  <!-- Non error -->

<div class="h-100 flex justify-center">
  <div v-if="loading" class="flex p-5 h-screen align-center justify-center">
      <semipolar-spinner :animation-duration="2000" :size="50" color="#1eb980" />
  </div>
<div v-else class="p-5 bg-foreground shadow rounded-xl">
    <ol-map :loadTilesWhileAnimating="true" :loadTilesWhileInteracting="true" style="height: 50rem; width: 50rem">

      <ol-zoom-control />

    <ol-view ref="view" :center="center" :rotation="rotation" :zoom="zoom" :projection="projection"/>

    <ol-tile-layer>
        <ol-source-osm />
    </ol-tile-layer>

          <ol-vector-layer>
            <ol-source-vector>
                <ol-feature>
                    <ol-geom-multi-point :coordinates="coordinates"></ol-geom-multi-point>
                    <ol-style>
                        <ol-style-circle :radius="radius">
                            <ol-style-fill :color="fillColor"></ol-style-fill>
                            <ol-style-stroke :color="strokeColor" :width="strokeWidth"></ol-style-stroke>
                        </ol-style-circle>
                    </ol-style>
                </ol-feature>

            </ol-source-vector>

        </ol-vector-layer>

    </ol-map>
  </div>
</div>
</template>
