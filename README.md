# NoGo Web

## 前端编译

### 安装依赖

```shell
npm install
```

### 编译
```shell
npm run build
```

## 后端编译

你需要先安装GCC（本人在7.5.0下编译通过）以及Go（1.15.2）

```shell
cd backend
go build
```

## 连接在一起

这里的工作目录还是backend文件夹

```shell
cp -r ../dist ./ui
```

## 启动！

```shell
./nogos
```

