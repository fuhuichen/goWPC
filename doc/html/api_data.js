define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./html/main.js",
    "group": "C__Users_fhhui_go_src_goWPC_doc_html_main_js",
    "groupTitle": "C__Users_fhhui_go_src_goWPC_doc_html_main_js",
    "name": ""
  },
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
            "description": "<p>最多回傳用戶數目</p>"
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
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "userList",
            "description": "<p>回傳用戶資訊陣列</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.id",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.email",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.name",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.mobile",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.imageUrl",
            "description": "<p>用戶圖片url</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "userList.registered",
            "description": "<p>是否已報到</p>"
          }
        ]
      }
    },
    "filename": "./src/example.js",
    "groupTitle": "Face_Recognition"
  },
  {
    "type": "post",
    "url": "/api/user/delete",
    "title": "刪除用戶",
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
    "filename": "./src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/info",
    "title": "尋找用戶",
    "version": "0.0.1",
    "name": "GetUser",
    "group": "User",
    "description": "<p>使用用戶的ID,email或mobile其中之一來取得用戶資訊，參數三種至少要有一種</p>",
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
            "field": "email",
            "description": "<p>用戶 email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "mobile",
            "description": "<p>用戶 mobile</p>"
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
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.name",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.email",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.company",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.mobile",
            "description": ""
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
            "field": "user.frEnabled",
            "description": "<p>是否已註冊照片</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "user.imageUrl",
            "description": "<p>用戶圖片url</p>"
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
            "description": "<p>簽到紀錄-簽到時間unix time</p>"
          }
        ]
      }
    },
    "filename": "./src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/create",
    "title": "建立用戶",
    "version": "0.0.1",
    "name": "PostUser",
    "group": "User",
    "description": "<p>建立新用戶, company+firstname+lastname是唯一</p>",
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
            "description": "<p>用戶公司名稱</p>"
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
    "filename": "./src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/update",
    "title": "更新用戶資訊",
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
            "description": "<p>用戶 ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>用戶姓名</p>"
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
            "field": "company",
            "description": "<p>用戶公司名稱</p>"
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
            "field": "registered",
            "description": "<p>用戶是否已經簽到</p>"
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
    "filename": "./src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/check",
    "title": "攤位簽到",
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
    "filename": "./src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/updateImage",
    "title": "更新用戶圖片",
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
    "filename": "./src/example.js",
    "groupTitle": "User"
  },
  {
    "type": "post",
    "url": "/api/user/list",
    "title": "列出使用者",
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
            "description": "<p>回傳用戶資訊陣列</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.id",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.email",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.name",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.mobile",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userList.imageUrl",
            "description": "<p>用戶圖片url</p>"
          },
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "userList.registered",
            "description": "<p>是否已報到</p>"
          }
        ]
      }
    },
    "filename": "./src/example.js",
    "groupTitle": "User"
  }
] });