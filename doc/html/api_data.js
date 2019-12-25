define({ "api": [
  {
    "type": "post",
    "url": "/api/fr/verification",
    "title": "人臉辨識",
    "version": "0.0.1",
    "name": "faceRecognition",
    "group": "Face_Recognition",
    "description": "<p>人臉辦識，可以選擇可信度，及最大回傳數目，結果會回傳0-max個辨識結果高於可信度值的用戶資訊</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "image",
            "description": "<p>臉部圖片(base64-encoded)</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "threshold",
            "description": "<p>可信度(0.0-1.0),數字愈高可信度愈高</p>"
          },
          {
            "group": "Parameter",
            "type": "Numnber",
            "optional": false,
            "field": "max",
            "description": "<p>最多回傳用戶數目 (目前只支援1個)</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功),userList若為空則代表沒有辨識到符合的臉部 1:INVALID_PARAMETERS(參數缺少或錯誤 4:INVALID_IMAGE(圖片格式不符)</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "userList",
            "description": "<p>回傳用戶資訊陣列(物件參考/api/user/info)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "Face_Recognition"
  },
  {
    "type": "post",
    "url": "/api/order/create",
    "title": "1.建立點餐資訊",
    "version": "0.0.1",
    "name": "CreateOrder",
    "group": "Order",
    "description": "<p>建立點餐資訊，使用者userId需要為資料庫中之用戶Id</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "userId",
            "description": "<p>使用者ID</p>"
          },
          {
            "group": "Parameter",
            "type": "Object[]",
            "optional": false,
            "field": "orderList",
            "description": "<p>點餐列表</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "orderList.name",
            "description": "<p>餐點名稱</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "orderList.amount",
            "description": "<p>餐點數量</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "orderList.ice",
            "description": "<p>餐點冰度</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "orderList.sugar",
            "description": "<p>餐點甜度</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "orderList.size",
            "description": "<p>餐點大小</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "orderList.padding",
            "description": "<p>附加品項</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤) 5:OPERATION_FAIL(建立失敗)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "orderNumber",
            "description": "<p>點餐代碼,四碼(例如0001)，每天(或超過9999)重置(若是成功建立)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "Order"
  },
  {
    "type": "post",
    "url": "/api/order/last",
    "title": "2.取得最後點餐資訊",
    "version": "0.0.1",
    "name": "FindLastOrder",
    "group": "Order",
    "description": "<p>取得使用者(userId)最後點餐記錄</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "userId",
            "description": "<p>使用者ID</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤) 6:NO_ORDER_EXIST(用戶沒有點過餐)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "order",
            "description": "<p>點餐資訊</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "orderNumber",
            "description": "<p>點餐代碼, 四碼(例如0001)，每天(或超過9999)重置</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "time",
            "description": "<p>點餐時間(UNIX TIME)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderList",
            "description": "<p>點餐資訊</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderList.name",
            "description": "<p>餐點名稱</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "order.orderList.amount",
            "description": "<p>餐點數量</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderList.ice",
            "description": "<p>餐點冰度</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderList.sugar",
            "description": "<p>餐點甜度</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderList.size",
            "description": "<p>餐點大小</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderList.padding",
            "description": "<p>附加品項</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "order.orderUser",
            "description": "<p>用戶資訊</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderUser.id",
            "description": "<p>用戶 ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderUser.firstname",
            "description": "<p>用戶名</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderUser.lastname",
            "description": "<p>用戶姓</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderUser.email",
            "description": "<p>用戶email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderUser.company",
            "description": "<p>用戶公司名稱</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "order.orderUser.mobile",
            "description": "<p>用戶手機號碼</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "Order"
  },
  {
    "type": "post",
    "url": "/api/order/list",
    "title": "3.列出點餐記錄",
    "version": "0.0.1",
    "name": "ListOrder",
    "group": "Order",
    "description": "<p>取得起始時間之後的所有點餐紀錄</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "time",
            "description": "<p>起始時間(UNIX TIMESTAMP)</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "orders",
            "description": "<p>點餐資訊,參考/api/order/last</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "Order"
  },
  {
    "type": "post",
    "url": "/api/order/isSpecialBonus",
    "title": "4.是否取得今日獎勵",
    "version": "0.0.1",
    "name": "isSpecialBonus",
    "group": "Order",
    "description": "<p>取得使用者(userId)是否領過今日獎勵</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "userId",
            "description": "<p>使用者ID</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "isSpecialBonus",
            "description": "<p>是否領過今日獎勵</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "Order"
  },
  {
    "type": "post",
    "url": "/api/order/setSpecialBonus",
    "title": "5.取得今日獎勵",
    "version": "0.0.1",
    "name": "setSpecialBonus",
    "group": "Order",
    "description": "<p>將使用者(userId)設定成已領過今日獎勵</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "userId",
            "description": "<p>使用者ID</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "Order"
  },
  {
    "type": "post",
    "url": "/api/user/delete",
    "title": "8.刪除用戶",
    "version": "0.0.1",
    "name": "DeleteUser",
    "group": "User",
    "description": "<p>刪除用戶資料</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>用戶 ID</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤) 3:USER_NOT_EXIST(用戶不存在)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/info",
    "title": "6.尋找用戶",
    "version": "0.0.1",
    "name": "GetUser",
    "group": "User",
    "description": "<p>使用用戶的ID來取得用戶資訊</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>用戶 ID</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤) 3:USER_NOT_EXIST(用戶不存在)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>用戶資訊</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.id",
            "description": "<p>用戶 ID</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "firstname",
            "description": "<p>用戶名</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "lastname",
            "description": "<p>用戶姓</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>用戶email</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "company",
            "description": "<p>用戶公司名稱</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "mobile",
            "description": "<p>用戶手機號碼</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "user.registered",
            "description": "<p>是否已報到</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "user.counterRegistered",
            "description": "<p>是否由櫃台報到</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "user.registerTime",
            "description": "<p>報到時間(UNIX TIMESTAMP)</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "user.faceRegistered",
            "description": "<p>是否已註冊照片</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "user.checkList",
            "description": "<p>用戶到攤位簽到紀錄</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.extend1",
            "description": "<p>自訂資料1</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.extend2",
            "description": "<p>自訂資料2</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.checkList.boothName",
            "description": "<p>簽到紀錄-攤位名稱</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "user.checkList.checked",
            "description": "<p>簽到紀錄-是否已簽到</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "user.checkList.time",
            "description": "<p>簽到紀錄-簽到時間(UNIX TIMESTAMP)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/face",
    "title": "5.取得用戶圖片",
    "version": "0.0.1",
    "name": "GetUserFace",
    "group": "User",
    "description": "<p>使用用戶的ID來取得用戶圖片</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>用戶 ID</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤) 3:USER_NOT_EXIST(用戶不存在)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "image",
            "description": "<p>用戶圖片(base64)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/create",
    "title": "1.建立用戶",
    "version": "0.0.1",
    "name": "PostUser",
    "group": "User",
    "description": "<p>建立新用戶, company+firstname+lastname+email是唯一</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "firstname",
            "description": "<p>用戶名(必填)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "lastname",
            "description": "<p>用戶姓(必填)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>用戶email(必填)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "company",
            "description": "<p>用戶公司名稱(必填)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>用戶公司職稱</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "mobile",
            "description": "<p>用戶手機號碼</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "extend1",
            "description": "<p>自訂資料1</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "extend2",
            "description": "<p>自訂資料2</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤) 2:USER_EXIST(用戶已存在) 5:OPERATION_FAIL(建立失敗)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>錯誤訊息</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>用戶ID (若是成功建立)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/update",
    "title": "2.更新用戶資訊",
    "version": "0.0.1",
    "name": "UpdateUser",
    "group": "User",
    "description": "<p>更新用戶資訊，包含基本資訊及用戶簽到與否</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>用戶 ID (必填)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>用戶email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "mobile",
            "description": "<p>用戶手機號碼</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>用戶職稱</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "extend1",
            "description": "<p>自訂資料1</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "extend2",
            "description": "<p>自訂資料2</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "registered",
            "description": "<p>用戶是否已經簽到</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "counterRegistered",
            "description": "<p>用戶是否由櫃台簽到</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤) 3:USER_NOT_EXIST(用戶不存在)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/updateCheck",
    "title": "4.攤位簽到",
    "version": "0.0.1",
    "name": "UpdateUserCheck",
    "group": "User",
    "description": "<p>若人臉辨識成功，在攤位簽到</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>用戶 ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "boothName",
            "description": "<p>攤位名稱</p>"
          },
          {
            "group": "Parameter",
            "type": "Boolean",
            "optional": false,
            "field": "checked",
            "description": "<p>是否已簽到</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤 3:USER_NOT_EXIST(用戶不存在)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/updateImage",
    "title": "3.更新用戶圖片",
    "version": "0.0.1",
    "name": "UpdateUserImage",
    "group": "User",
    "description": "<p>更新用戶圖片，若格式不符會回傳錯誤 4:INVALID_IMAGE(圖片格式不符)</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>用戶 ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "image",
            "description": "<p>用戶臉部圖片(base64-encoded)</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代碼 0:SUCCESS(成功) 1:INVALID_PARAMETERS(參數缺少或錯誤 3:USER_NOT_EXIST(用戶不存在) 4:INVALID_IMAGE(圖片格式不符)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/list",
    "title": "7.列出使用者",
    "version": "0.0.1",
    "name": "listUsers",
    "group": "User",
    "description": "<p>可以透過關鍵字搜索名字/公司名稱及Email，列出含有字串的使用者,不使用關鍵字則列出全部使用者</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "keyword",
            "description": "<p>關鍵字</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>錯誤代不使用關鍵字則列出全部使用者 0:SUCCESS(成功)</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "userList",
            "description": "<p>回傳用戶資訊陣列(物件參考/api/user/info)</p>"
          }
        ]
      }
    },
    "filename": "src/example.js",
    "groupTitle": "User"
  }
] });
