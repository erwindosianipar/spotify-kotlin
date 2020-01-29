package com.example.architecturecomponent.model

import com.google.gson.annotations.SerializedName

class Artist(name: String = "", debut: String = "", image: String = "", category: String = "") {

    @SerializedName("Name")
    var name: String = name

    @SerializedName("Debut")
    var debut: String = debut

    @SerializedName("Image")
    var image: String = image

    @SerializedName("Category")
    var category: String = category

    override fun toString(): String {
        return "Artist(name='$name', debut='$debut', category='$category')"
    }
}