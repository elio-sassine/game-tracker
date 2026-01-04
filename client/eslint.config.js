const eslint = require("@eslint/js");
const tseslint = require("typescript-eslint");
const angular = require("angular-eslint");

module.exports = tseslint.config(
    {
        // 1. GLOBAL SETTINGS
        files: ["**/*.ts"],
        extends: [
            eslint.configs.recommended,
            ...tseslint.configs.recommended,
            ...tseslint.configs.stylistic,
            ...angular.configs.tsRecommended,
        ],
        processor: angular.processInlineTemplates,
        rules: {
            // Angular specific rules
            "@angular-eslint/directive-selector": [
                "error",
                {
                    type: "attribute",
                    prefix: "app",
                    style: "camelCase",
                },
            ],
            "@angular-eslint/component-selector": [
                "error",
                {
                    type: "element",
                    prefix: "app",
                    style: "kebab-case",
                },
            ],
            // Common customizations
            "@typescript-eslint/no-explicit-any": "warn", // Warn instead of error for 'any'
            "no-console": ["warn", { allow: ["warn", "error"] }], // Allow console.warn/error
        },
    },
    {
        // 2. TEMPLATE (HTML) SETTINGS
        files: ["**/*.html"],
        extends: [
            ...angular.configs.templateRecommended,
            ...angular.configs.templateAccessibility,
        ],
        rules: {
            // Accessibility rules (optional but recommended)
            "@angular-eslint/template/click-events-have-key-events": "warn",
            "@angular-eslint/template/interactive-supports-focus": "warn",
        },
    }
);
