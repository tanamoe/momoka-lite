# Metadata

Working with metadata returned from [`titles`-related collection](https://github.com/tanamoe/momoka-lite/blob/master/hooks/image_secret.go#L26).

## Images
Images returned from metadata can be use as source with [image endpoint](/getting-started#images).

An example with native `<img />` tag:
```html
<img loading="lazy"
     src="https://image.tana.moe/WpkWxw9w_jy4hjV1y7Bep28sT0fpxdziNMVe5FTe/1920x0/filters:quality(90)/publications/7n8sxmp2dwbcx1e/my_vi_ham_nguc_tap_12_RxSSJtewGk.jpg"
     srcset="https://image.tana.moe/vNAF44jG49H_uzXUSCua4xHCMGNOilp4ktepjd9U/1280x0/filters:quality(90)/publications/7n8sxmp2dwbcx1e/my_vi_ham_nguc_tap_12_RxSSJtewGk.jpg 1280w,
             https://image.tana.moe/K5KBiJDYBPhfdRoKoC6uupiA2fs5RN7KDCoQ5ZWf/120x0/filters:quality(90)/publications/7n8sxmp2dwbcx1e/my_vi_ham_nguc_tap_12_RxSSJtewGk.jpg 160w,
             https://image.tana.moe/WpkWxw9w_jy4hjV1y7Bep28sT0fpxdziNMVe5FTe/1920x0/filters:quality(90)/publications/7n8sxmp2dwbcx1e/my_vi_ham_nguc_tap_12_RxSSJtewGk.jpg 1920w,
             https://image.tana.moe/L06t6wTZlXFc6abN6AImYxIwUbpYzLaE3rnMuVoN/320x0/filters:quality(90)/publications/7n8sxmp2dwbcx1e/my_vi_ham_nguc_tap_12_RxSSJtewGk.jpg 320w,
             https://image.tana.moe/8tfzDc0B2w4dhYUG1wAVstyxx4M860DfNx2UchNm/480x0/filters:quality(90)/publications/7n8sxmp2dwbcx1e/my_vi_ham_nguc_tap_12_RxSSJtewGk.jpg 480w,
             https://image.tana.moe/Y_MnKGORxtftn8H19oErI0l1O_Cpg8UNFYs_6zE-/640x0/filters:quality(90)/publications/7n8sxmp2dwbcx1e/my_vi_ham_nguc_tap_12_RxSSJtewGk.jpg 640w"
     alt="Mỹ vị hầm ngục - Tập 12"
     sizes="(max-width: 640px) 40vw, (max-width: 768px) 30vw, 20vw"
/>
```

## Schema
| Field | Type | Notes |
|---|---|---|
| images | [MetadataImages \| MetadataImages\[\]](https://github.com/tanamoe/kikuri/blob/410e10d4cdd143ad212a9bf3176f9ff8a2ccb363/types/common.ts#L8C13-L8C27) | [imagor](https://github.com/cshum/imagor#image-endpoint) image URI |

## Response
::: code-group
```zsh [Request]
GET /api/collections/publications/records/gqw2iu33ucsd8pu
```
```json [Response]
{
  "metadata": {
    "images": [
      {
        "1280w": "Uz86rGbjygq2isQTtajP0Rtr-r8lTqTBYqTtaOl_/1280x0/filters:quality(90)/guv9vnyfu5pdz9t/gqw2iu33ucsd8pu/toi_yeu_nu_phan_dien_tap_3_FUdZPXjXKn.jpg",
        "160w": "EODzD_6Ymdh7aYLzrBe85qpOJO1qLZjcC6rFh6i5/120x0/filters:quality(90)/guv9vnyfu5pdz9t/gqw2iu33ucsd8pu/toi_yeu_nu_phan_dien_tap_3_FUdZPXjXKn.jpg",
        "1920w": "Ny3sp5XuuIhqgG1n0xs_64smEpIORMvFsGPlZcft/1920x0/filters:quality(90)/guv9vnyfu5pdz9t/gqw2iu33ucsd8pu/toi_yeu_nu_phan_dien_tap_3_FUdZPXjXKn.jpg",
        "320w": "G6IWnyBu8tdwOtAj043yrkYDlYXmPs5BxC5yGb_q/320x0/filters:quality(90)/guv9vnyfu5pdz9t/gqw2iu33ucsd8pu/toi_yeu_nu_phan_dien_tap_3_FUdZPXjXKn.jpg",
        "480w": "vssLb5ZtTpuwH-QAtVHMgxOVOVAanxMtpSvl6SpQ/480x0/filters:quality(90)/guv9vnyfu5pdz9t/gqw2iu33ucsd8pu/toi_yeu_nu_phan_dien_tap_3_FUdZPXjXKn.jpg",
        "640w": "SbqHxAUAYNSp_T8pr3RhD-2sPjBDpznJVi_BrIvD/640x0/filters:quality(90)/guv9vnyfu5pdz9t/gqw2iu33ucsd8pu/toi_yeu_nu_phan_dien_tap_3_FUdZPXjXKn.jpg"
      }
    ]
  }
  // ...
}
```
:::
