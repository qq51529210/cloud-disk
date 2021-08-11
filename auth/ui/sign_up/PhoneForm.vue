<template>
    <var-form ref="form" :disabled="disabled">
        <!-- number input -->
        <var-input
            :placeholder="t('placeholderNumber')"
            v-model="formData.number"
            clearable
            :rules="rulePhoneNumber"
        />
        <!-- password input -->
        <var-input
            :type="showPassword ? 'text' : 'password'"
            :placeholder="t('placeholderPassword')"
            v-model="formData.password"
            clearable
        />
        <var-rate v-model="passwordScore" :size="14" />
        <var-row justify="space-between" align="flex-end" :gutter="10">
            <!-- verification code input -->
            <var-col :span="16">
                <var-input
                    :placeholder="t('placeholderVerificationCode')"
                    v-model="formData.code"
                    clearable
                    :rules="ruleVerificationCode"
                />
            </var-col>
            <!-- send sms button -->
            <var-col :span="8">
                <var-row justify="flex-end" align="flex-end">
                    <var-button
                        type="primary"
                        :disabled="smsButtonDisabled || disabled"
                        block
                        @click="onSendSMS"
                    >
                        <template #default>
                            <var-countdown
                                v-if="smsButtonDisabled"
                                :time="smsButtonCounter"
                                format="ss"
                                @end="smsButtonDisabled = false"
                            ></var-countdown>
                            <span v-else>{{ t("sendSMS") }}</span>
                        </template>
                    </var-button>
                </var-row>
            </var-col>
        </var-row>
        <var-button
            type="primary"
            block
            size="large"
            @click="onSignUp"
            :loading="disabled"
            :disabled="disabled"
        >{{ t("signUp") }}</var-button>
    </var-form>
</template>

<script setup>
// vue
import { ref, reactive, computed } from 'vue'
// i18n
import { useI18n } from "vue-i18n";
const { t } = useI18n()
// Disabled form
const disabled = ref(false)
// rules
import isMobilePhone from "validator/es/lib/isMobilePhone"
import isEmail from "validator/es/lib/isEmail"
import isNumeric from "validator/es/lib/isNumeric"
const rulePhoneNumber = [v => isMobilePhone(v) || t("errPhoneNumber")]
const ruleEmail = [v => isEmail(v) || t("errEmail")]
const ruleVerificationCode = [
    v => (isNumeric(v) && v.length === 6) || t("errVerificationCode")
]
// Form
const form = ref(null)
const formData = reactive({
    number: "",
    password: "",
    code: "",
})
const passwordScore = ref(0)
// Send SMS button
const smsButtonCounter = ref(0)
const smsButtonDisabled = computed({
    get: () => smsButtonCounter.value > 0,
    set: (value) => {
        if (!value) {
            smsButtonCounter.value = 0
        }
    },
})
const onSendSMS = () => {
    smsButtonCounter.value = 10000
}
// Handle sign up
const onSignUp = () => {
    disabled.value = true;
    phoneForm.value
        .validate()
        .then((v) => {
            disabled.value = true;
            console.log(v);
        })
        .catch((e) => {
            disabled.value = true;
            console.log(v);
        });
}
</script>