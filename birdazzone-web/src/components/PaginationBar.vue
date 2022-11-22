<scritp setup lang="ts">
    import {ref} from 'vue'
    const props = defineProps<{actualPage:number, max:number}>()
    
    let tabs = []
    for(let i = 1; i<=props.max;i++){
        tabs +={index:i, actual:(props.actualPage==i)?true:false,}
    }
    
    /* la mia idea sarebbe creare un oggetto del tipo {index: number, actual: boolean}
        poi fare un ciclo per max volte in cui si popola tabs, ovvero la lista di pulsanti
        in questo modo si riesce a controllare la barra dei numeri e a sapere che sfondo mettere
    */
</scritp>

<template>
    <div class="flex flex-1 flex-row justify-center">
        <div id="firstIndex" class="mx-2 p-1 text-lgreen text-center hover:bg-lgreen hover:text-foreground" style="flex: 1 1 auto; maxHeight:2.2rem; maxWidth:2.2rem;border-radius:50%; border:solid 2px;">
            {{tabs[0].index}}
        </div>
        
        <div v-if="actualPage - 2 > 1">
            ...
        </div>

        <div v-else >

        </div>

        <div v-if="actualPage + 2 < max">
            ...
        </div>
        <div id="firstIndex" class="mx-2 p-1 text-lgreen text-center hover:bg-lgreen hover:text-foreground" style="flex: 1 1 auto; maxHeight:2.2rem; maxWidth:2.2rem;border-radius:50%; border:solid 2px;">
            {{props.max}}
        </div>
    </div>
</template>