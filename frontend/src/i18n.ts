import { createI18n } from 'vue-i18n'

const messages = {
  fr: {
    errors: {
      ERR_EMPTY_JID: "Le JID est requis.",
      ERR_INVALID_JID: "Le JID doit contenir un '@'.",
      ERR_PASSWORD_REQUIRED: "Le mot de passe est requis.",
      ERR_ACCOUNT_EXISTS: "Un compte avec ce JID est déjà connecté.",
      ERR_UNKNOWN: "Une erreur inconnue est survenue."
    },
    account: {
      added_success: "Compte ajouté avec succès !"
    }
  },
  en: {
    errors: {
      ERR_EMPTY_JID: "JID is required.",
      ERR_INVALID_JID: "JID must contain '@'.",
      ERR_PASSWORD_REQUIRED: "Password is required.",
      ERR_ACCOUNT_EXISTS: "An account with this JID is already connected.",
      ERR_UNKNOWN: "An unknown error occurred."
    },
    account: {
      added_success: "Account added successfully!"
    }
  }
}

const i18n = createI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  messages,
})

export default i18n
