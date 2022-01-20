<script setup>
import QRCodeVue3 from 'qrcode-vue3'
import FlexBoxVue from '../components/FlexBox.vue';

const size = 200
const backgroundOptions = { color: '#ffffff' }
const qrOptions = { typeNumber: 0, mode: 'Byte', errorCorrectionLevel: 'H' }
const imageOptions = { hideBackgroundDots: true, imageSize: 0.4, margin: 0 }
const dotsOptions = {
    type: 'extra-rounded',
    color: '#26249a',
    gradient: {
        type: 'linear',
        rotation: 0,
        colorStops: [
            { offset: 0, color: '#26249a' },
            { offset: 1, color: '#26249a' },
        ],
    },
}
const cornersSquareOptions = { type: 'square', color: '#000000' }
const cornersDotOptions = { type: undefined, color: '#000000' }
</script>

<template>
    <FlexBoxVue justify="center">
        <QRCodeVue3
            value="test qrcode"
            :width="size"
            :height="size"
            :backgroundOptions="backgroundOptions"
            :qrOptions="qrOptions"
            :imageOptions="imageOptions"
            :dotsOptions="dotsOptions"
            :cornersSquareOptions="cornersSquareOptions"
            :cornersDotOptions="cornersDotOptions"
        />
    </FlexBoxVue>
</template>
