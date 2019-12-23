/**
@apiDefine  CreateUserError

**/


/**
 * @api {post} /api/user/create 1.建立用戶
 * @apiVersion 0.0.1
 * @apiName PostUser
 * @apiGroup User
 *
 * @apiDescription 建立新用戶, company+firstname+lastname+email是唯一
 *
 * @apiParam {String}   firstname    用戶名(必填)
 * @apiParam {String}   lastname     用戶姓(必填)
 * @apiParam {String}   email       用戶email(必填)
 * @apiParam {String}   company     用戶公司名稱(必填)
 * @apiParam {String}   title       用戶公司職稱
 * @apiParam {String}   mobile      用戶手機號碼
 * @apiParam {String}   extend1     自訂資料1
 * @apiParam  {String}  extend2     自訂資料2

 * @apiSuccess {Number} code  錯誤代碼
 *                 0:SUCCESS(成功)
 *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
 *                 2:USER_EXIST(用戶已存在)
 *                 5:OPERATION_FAIL(建立失敗)
 * @apiSuccess {String} message  錯誤訊息
 * @apiSuccess {String} id   用戶ID (若是成功建立)
 *
 */
 function postUser() { return; }

 /**
  * @api {post} /api/user/info  6.尋找用戶
  * @apiVersion 0.0.1
  * @apiName GetUser
  * @apiGroup User
  *
  * @apiDescription 使用用戶的ID來取得用戶資訊
  *
  * @apiParam {String} id  用戶 ID
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
  *                 3:USER_NOT_EXIST(用戶不存在)
  * @apiSuccess {String}   message  錯誤訊息
  * @apiSuccess  {Object}  user         用戶資訊
  * @apiSuccess  {String}    user.id    用戶 ID
  * @apiSuccess {String}   firstname    用戶名
  * @apiSuccess {String}   lastname     用戶姓
  * @apiSuccess {String}   email       用戶email
  * @apiSuccess {String}   company     用戶公司名稱
  * @apiSuccess  {String}   mobile      用戶手機號碼
  * @apiSuccess   {Boolean}  user.registered 是否已報到
  * @apiSuccess   {Boolean}  user.counterRegistered 是否由櫃台報到
  * @apiSuccess   {Boolean}  user.registerTime    報到時間
  * @apiSuccess   {Boolean}  user.faceRegistered 是否已註冊照片
  * @apiSuccess   {Object[]} user.checkList     用戶到攤位簽到紀錄
  * @apiSuccess   {String}   user.extend1     自訂資料1
  * @apiSuccess    {String}  user.extend2     自訂資料2
  * @apiSuccess   {String}   user.checkList.boothName  簽到紀錄-攤位名稱
  * @apiSuccess   {Boolean}  user.checkList.checked  簽到紀錄-是否已簽到
  * @apiSuccess   {Number}   user.checkList.time 簽到紀錄-簽到時間unix time
  *
  */

 function getUser() { return; }
 /**
  * @api {post} /api/user/face  5.取得用戶圖片
  * @apiVersion 0.0.1
  * @apiName GetUserFace
  * @apiGroup User
  *
  * @apiDescription 使用用戶的ID來取得用戶圖片
  *
  * @apiParam {String} id  用戶 ID
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
  *                 3:USER_NOT_EXIST(用戶不存在)
  * @apiSuccess {String}   message  錯誤訊息
  * @apiSuccess  {String}  image         用戶圖片(base64)
  *
  */

 function getUserFace() { return; }

 /**
  * @api {post} /api/user/updateCheck  4.攤位簽到
  * @apiVersion 0.0.1
  * @apiName UpdateUserCheck
  * @apiGroup User
  *
  * @apiDescription 若人臉辨識成功，在攤位簽到
  *
  * @apiParam {String}     id        用戶 ID
  * @apiParam {String}     boothName 攤位名稱
  * @apiParam {Boolean}    checked  是否已簽到
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *
  */
  function updateUserCheck() { return; }

 /**
  * @api {post} /api/user/updateImage  3.更新用戶圖片
  * @apiVersion 0.0.1
  * @apiName UpdateUserImage
  * @apiGroup User
  *
  * @apiDescription 更新用戶圖片，若格式不符會回傳錯誤 4:INVALID_IMAGE(圖片格式不符)
  *
  * @apiParam {String} id  用戶 ID
  * @apiParam {String} image 用戶臉部圖片(base64-encoded)
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *                 4:INVALID_IMAGE(圖片格式不符)
  *
  */
function updateUserCheck() { return; }

  /**
   * @api {post} /api/user/update  2.更新用戶資訊
   * @apiVersion 0.0.1
   * @apiName UpdateUser
   * @apiGroup User
   *
   * @apiDescription 更新用戶資訊，包含基本資訊及用戶簽到與否
   *
   * @apiParam {String} id  用戶 ID (必填)
   * @apiParam {String} email 用戶email
   * @apiParam {String} mobile  用戶手機號碼
   * @apiParam {String} title  用戶職稱
   * @apiParam {String}  extend1     自訂資料1
   * @apiParam  {String}  extend2     自訂資料2
   * @apiParam {Bool}     registered 用戶是否已經簽到
   * @apiParam {Bool}     counterRegistered 用戶是否由櫃台簽到
   *
   * @apiSuccess {Number} code  錯誤代碼
   *                 0:SUCCESS(成功)
   *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
   *                 3:USER_NOT_EXIST(用戶不存在)
   *
   */
  function updateUser() { return; }

  /**
   * @api {post} /api/user/delete  8.刪除用戶
   * @apiVersion 0.0.1
   * @apiName DeleteUser
   * @apiGroup User
   *
   * @apiDescription 刪除用戶資料
   *
   * @apiParam {String} id  用戶 ID
   *
   * @apiSuccess {Number} code  錯誤代碼
   *                 0:SUCCESS(成功)
   *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
   *                 3:USER_NOT_EXIST(用戶不存在)
   *
   */

  function deleteUser() { return; }

  /**
   * @api {post} /api/user/list 7.列出使用者
   * @apiVersion 0.0.1
   * @apiName listUsers
   * @apiGroup User
   * @apiDescription 可以透過關鍵字搜索名字/公司名稱及Email，列出含有字串的使用者,不使用關鍵字則列出全部使用者
   * @apiParam {String} keyword 關鍵字
   *
   * @apiSuccess {Number} code  錯誤代不使用關鍵字則列出全部使用者
   *                 0:SUCCESS(成功)
   * @apiSuccess {Object[]} userList  回傳用戶資訊陣列(物件參考/api/user/info)
   */
   function listUers() { return; }
 /**
  * @api {post} /api/fr/verification 人臉辨識
  * @apiVersion 0.0.1
  * @apiName faceRecognition
  * @apiGroup Face Recognition
  * @apiDescription 人臉辦識，可以選擇可信度，及最大回傳數目，結果會回傳0-max個辨識結果高於可信度值的用戶資訊
  * @apiParam {String} image 臉部圖片(base64-encoded)
  * @apiParam {Number}  threshold  可信度(0.0-1.0),數字愈高可信度愈高
  * @apiParam {Numnber} max  最多回傳用戶數目 (目前只支援1個)
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *                 4:INVALID_IMAGE(圖片格式不符)
  * @apiSuccess {Object[]} userList 回傳用戶資訊陣列(物件參考/api/user/info)
  */
  function faceRecognition() { return; }
  /*
  type OrderItem struct {
  	Name          string         			`json:"name" bson:"name"`
  	Amount        int64								`json:"amount" bson:"amount"`
  	Size          string         			`json:"size" bson:"size"`
  	Ice           string         			`json:"ice" bson:"ice"`
  	Sugar         string         		  `json:"sugar" bson:"sugar"`
  	Padding       string           		`json:"padding" bson:"padding"`
  }


  type OrderUser struct {
  	FirstName    string         `json:"firstname" bson:"firstname"`
  	LastName     string         `json:"lastname" bson:"lastname"`
  	Email        string         `json:"email"  bson:"email"`
  	Company      string         `json:"company" bson:"company"`
  	Title        string     	  `json:"title" bson:"title"`
  	Mobile       string         `json:"mobile" bson:"mobile"`
  	Extend1      string 			  `json:"extend1" bson:"extend1"`
  	Extend2      string 				`json:"extend2" bson:"extend2"`
  }

  type Order struct {
  	ID           		   bson.ObjectId  `json:"id" bson:"_id,omitempty"`
  	OrderNumber        int64         	`json:"orderNumber" bson:"orderNumber"`
  	UserId    	       string         `json:"userid" bson:"userid"`
  	Time               int64          `json:"time" bson:"time"`
  	UserInfo           OrderUser			`json:"orderUser" bson:"orderUser"`
  	List          		 []OrderItem 	  `json:"orderList" bson:"checkUser"`
  }
  */

  /**
   * @api {post} /api/order/create 1.建立點餐資訊
   * @apiVersion 0.0.1
   * @apiName CreateOrder
   * @apiGroup Order
   *
   * @apiDescription 建立點餐資訊，使用者UserId需要為資料庫中之用戶Id
   *
   * @apiParam     {String}   UserId          使用者ID
   * @apiParam     {Object[]}   orderList        點餐列表
   * @apiParam     {String}  orderList.name      餐點名稱
   * @apiParam     {Number}  orderList.amount    餐點數量
   * @apiParam     {String}  orderList.ice       餐點冰度
   * @apiParam     {String}  orderList.sugar     餐點甜度
   * @apiParam     {String}  orderList.size      餐點大小
   * @apiParam     {String}  orderList.padding   附加品項

   * @apiSuccess {Number} code  錯誤代碼
   *                 0:SUCCESS(成功)
   *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
   *                 5:OPERATION_FAIL(建立失敗)
   * @apiSuccess {String} message  錯誤訊息
   * @apiSuccess {String} orderNumber   點餐代碼 (若是成功建立)
   *
   */
   function CreateOrder() { return; }

   /**
    * @api {post} /api/order/last 2.取得最後點餐資訊
    * @apiVersion 0.0.1
    * @apiName FindLastOrder
    * @apiGroup Order
    *
    * @apiDescription 取得使用者(UserId)最後點餐記錄
    *
    * @apiParam     {String}     UserId          使用者ID
    * @apiSuccess {Number} code  錯誤代碼
    *                 0:SUCCESS(成功)
    *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
    *                 6:NO_ORDER_EXIST(用戶沒有點過餐)
    * @apiSuccess     {String}  message  錯誤訊息
    * @apiSuccess     {Object}  order                    點餐資訊
    * @apiSuccess     {String}  orderNumber              點餐代碼
    * @apiSuccess     {Number}  time                     點餐時間(UNIX TIME)
    * @apiSuccess     {String}  order.orderList           點餐資訊
    * @apiSuccess     {String}  order.orderList.name      餐點名稱
    * @apiSuccess     {Number}  order.orderList.amount    餐點數量
    * @apiSuccess     {String}  order.orderList.ice       餐點冰度
    * @apiSuccess     {String}  order.orderList.sugar     餐點甜度
    * @apiSuccess     {String}  order.orderList.size      餐點大小
    * @apiSuccess     {String}  order.orderList.padding   附加品項
    * @apiSuccess     {Object}  order.orderUser             用戶資訊
    * @apiSuccess     {String}  order.orderUser.id           用戶 ID
    * @apiSuccess     {String}  order.orderUser.firstname    用戶名
    * @apiSuccess     {String}  order.orderUser.lastname     用戶姓
    * @apiSuccess     {String}  order.orderUser.email       用戶email
    * @apiSuccess     {String}  order.orderUser.company     用戶公司名稱
    * @apiSuccess     {String}  order.orderUser.mobile      用戶手機號碼
    *
    */
    function FindLastOrder() { return; }


    /**
     * @api {post} /api/order/list  3.列出點餐記錄
     * @apiVersion 0.0.1
     * @apiName ListOrder
     * @apiGroup Order
     *
     * @apiDescription 取得起始時間之後的所有點餐紀錄
     *
     * @apiParam     {Number}     time          起始時間(UNIX TIMESTAMP)

     * @apiSuccess {Number} code  錯誤代碼
     *                 0:SUCCESS(成功)
     *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
     * @apiSuccess     {String}    message  錯誤訊息
     * @apiSuccess     {Object[]}  orders  點餐資訊,參考/api/order/last
     *
     */
     function FindLastOrder() { return; }

     /**
      * @api {post} /api/order/isSpecialBonus  4.是否取得今日獎勵
      * @apiVersion 0.0.1
      * @apiName isSpecialBonus
      * @apiGroup Order
      *
      * @apiDescription 取得使用者(UserId)是否領過今日獎勵
      *
      * @apiParam     {String}     UserId          使用者ID
      * @apiSuccess     {Number} code  錯誤代碼
      *                 0:SUCCESS(成功)
      *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
      * @apiSuccess     {String}  message  錯誤訊息
      * @apiSuccess     {Object}  isSpecialBonus           是否領過今日獎勵
      */
      function isSpecialBonus() { return;

        /**
         * @api {post} /api/order/setSpecialBonus  5.取得今日獎勵
         * @apiVersion 0.0.1
         * @apiName setSpecialBonus
         * @apiGroup Order
         *
         * @apiDescription 將使用者(UserId)設定成已領過今日獎勵
         *
         * @apiParam       {String}     UserId          使用者ID
         * @apiSuccess     {Number} code  錯誤代碼
         *                 0:SUCCESS(成功)
         *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
         * @apiSuccess     {String}  message  錯誤訊息
         */
         function setSpecialBonus() { return; }
