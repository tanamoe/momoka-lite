# titles

### <Badge type="info" text="GET" /> /api/collections/title/records

The **titles** collection serve book series.

## Schema
| Field | Type | Notes |
|---|---|---|
| id | string | |
| name | string | |
| description | string | html, should be sanitize |
| format | string | [formats](/reference/collections/formats) relation |
| cover | string | image path |
| genres | string[] | [genres](/reference/collections/genres) relation |
| created | timestamp | |
| updated | timestamp | |

## List/Search
::: code-group
```zsh [Request]
GET /api/collections/titles/records
```

```json [Response]
{
  "page": 1,
  "perPage": 30,
  "totalItems": 707,
  "totalPages": 24,
  "items": [
    {
      "collectionId": "s91oidzeo1xm4m7",
      "collectionName": "titles",
      "cover": "new_game_tap_1_cggq_w8d_xix_pgoGzgJK2K.png",
      "created": "2023-07-09 11:56:06.434Z",
      "description": "<p>Aoba Suzukaze là một học sinh tốt nghiệp trung học và là một nhà thiết kế nhân vật đầy tham vọng. Cô gia nhập công ty trò chơi điện tử Eagle Jump, được biết đến với trò chơi mà cô vô cùng yêu thích – Fairies Story.</p><p>Tại Eagle Jump, Aoba được xếp vào một team sáu người, bao gồm người thiết kế nhân vật chính cho Fairy Story, Kou Yagami. Được sự dìu dắt bởi hình mẫu của cô ấy, Aoba cố gắng trở thành một nhà thiết kế nhân vật giỏi hơn và tìm hiểu các kỹ năng của ngành công nghiệp trò chơi điện tử.</p><p>Đây là câu chuyện về những cô gái trên con đường theo đuổi ước mơ của mình.</p><p>(Nguồn: Mori Manga)</p>",
      "format": "tt6995wq46wqxkr",
      "genres": [
        "4uarrkkte6a9tso"
      ],
      "id": "5mjqcj8jfo2k263",
      "metadata": null,
      "name": "New Game",
      "updated": "2023-10-23 07:49:27.581Z"
    }
    // ...
  ]
}
```

:::

## View
::: code-group
```zsh [Request]
GET /api/collections/titles/records/[id]
```

```json [Response]
{
  "collectionId": "s91oidzeo1xm4m7",
  "collectionName": "titles",
  "cover": "bi_duoi_khoi_nhom_anh_hung_toi_muon_song_tu_do_tu_tai_o_vuong_do_TP2DqnCxVX.png",
  "created": "2023-10-22 19:56:34.873Z",
  "description": "<p>Flum Apricot chưa bao giờ có ý muốn rời khỏi ngôi làng của mình. Cô hài lòng dành cả quãng đời để sống yên bình tại nơi đây. Thế nhưng cuộc đời chẳng như mơ, sau khi bị chỉ định bởi lời tiên tri của Đấng Sáng tạo Origin rằng phải tham gia nhóm Anh hùng và đánh bại Ma Vương, cô không còn gì ngoài đau khổ. Làm sao một người có mọi chỉ số đều bằng 0 lại có tác dụng trong trận chiến?Kĩ năng duy nhất của cô là “Nghịch đảo” (Inverse), nhưng thậm chí cô còn không biết tác dụng hay ý nghĩa của nó. Hiền Nhân Jean Intage - một người đồng đội trong nhóm Anh hùng - luôn ngứa mắt với sự tồn tại của cô ở trong nhóm, đồng thời luôn nỗ lực để loại bỏ cô khỏi nhóm. Trong vực sâu của nỗi tuyệt vọng, khi lời nguyền cuối cùng cũng “đảo ngược” chính nó, Flum sẽ còn lại gì? Cô sẽ làm gì để thay đổi cuộc sống của mình?</p><p>(Nguồn: Ichi Light Novel)</p>",
  "format": "73hx8goiqg8kqjh",
  "genres": [
    "4uarrkkte6a9tso"
  ],
  "id": "v8vhf64w4xtwdhz",
  "metadata": null,
  "name": "Bị đuổi khỏi nhóm Anh hùng, tôi muốn sống tự do tự tại ở Vương đô",
  "updated": "2023-10-22 20:01:23.264Z"
}
```
:::
