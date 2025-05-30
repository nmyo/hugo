---
title: transform.Emojify 
description: Runs a string through the Emoji emoticons processor.
categories: []
keywords: []
params:
  functions_and_methods:
    aliases: [emojify]
    returnType: template.HTML
    signatures: [transform.Emojify INPUT]
aliases: [/functions/emojify]
---

`emojify` runs a passed string through the Emoji emoticons processor.

See the list of [emoji shortcodes] for available emoticons.

The `emojify` function can be called in your templates but not directly in your content files by default. For emojis in content files, set [`enableEmoji`] to `true` in your site's configuration. Then you can write emoji shorthand directly into your content files;

```text
I :heart: Hugo!
```

I :heart: Hugo!

[`enableEmoji`]: /configuration/all/#enableemoji
[emoji shortcodes]: /quick-reference/emojis/
