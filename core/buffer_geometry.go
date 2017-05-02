/*
 * Ported from three.js by @rydrman
 */

package core

import (
	"github.com/golang/glog"

	"github.com/rydrman/three.go/math3"
)

var BufferGeometryMaxIndex = 65535

type BufferGeometry struct {
	ID   int
	UUID string
	Name string
	Type string

	index           int
	Attributes      map[string]BufferAttribute
	MorphAttributes map[string]interface{}

	Groups []int

	BoundingBox   *math3.Box3
	BoudingSphere *math3.Sphere

	DrawRange float32

	*EventDispatcher
}

func (geo *BufferGeometry) getIndex() int {

	return geo.index

}

/*func (geo *BufferGeometry) setIndex( []index int) {

    geo.Index = index

	if ( Array.isArray( index ) ) {

		geo.Index = new ( arrayMax( index ) > 65535 ? Uint32BufferAttribute : Uint16BufferAttribute )( index, 1 );

	} else {

		geo.index = index;

	}

}*/

func (geo *BufferGeometry) AddAttribute(name string, attribute BufferAttribute) {

	if name == "index" {

		// TODO set back to warning when setIndex is ready
		glog.Fatalln("three.core.BufferGeometry.addAttribute: Use .setIndex() for index attribute.")
		//geo.SetIndex(attribute)

	}

	geo.Attributes[name] = attribute

}

/*

	func (geo *BufferGeometry) GetAttribute( name ) {

		return geo.Attributes[ name ];

	}

	func (geo *BufferGeometry) RemoveAttribute( name ) {

		delete geo.Attributes[ name ];

		return this;

	}

	func (geo *BufferGeometry) AddGroup( start, count, materialIndex ) {

		geo.Groups.push( {

			start: start,
			count: count,
			materialIndex: materialIndex !== undefined ? materialIndex : 0

		} );

	}

	func (geo *BufferGeometry) ClearGroups() {

		geo.Groups = [];

	}

	func (geo *BufferGeometry) SetDrawRange( start, count ) {

		geo.DrawRange.start = start;
		geo.DrawRange.count = count;

	}

	func (geo *BufferGeometry) ApplyMatrix( matrix ) {

		var position = geo.Attributes.position;

		if ( position !== undefined ) {

			matrix.applyToBufferAttribute( position );
			position.needsUpdate = true;

		}

		var normal = geo.Attributes.normal;

		if ( normal !== undefined ) {

			var normalMatrix = new Matrix3().getNormalMatrix( matrix );

			normalMatrix.applyToBufferAttribute( normal );
			normal.needsUpdate = true;

		}

		if ( geo.BoundingBox !== null ) {

			geo.ComputeBoundingBox();

		}

		if ( geo.BoundingSphere !== null ) {

			geo.ComputeBoundingSphere();

		}

		return this;

	}

	func (geo *BufferGeometry) RotateX() {

		// rotate geometry around world x-axis

		var m1;

		return function rotateX( angle ) {

			if ( m1 === undefined ) m1 = new Matrix4();

			m1.makeRotationX( angle );

			geo.ApplyMatrix( m1 );

			return this;

		};

	}(),

	func (geo *BufferGeometry) RotateY() {

		// rotate geometry around world y-axis

		var m1;

		return function rotateY( angle ) {

			if ( m1 === undefined ) m1 = new Matrix4();

			m1.makeRotationY( angle );

			geo.ApplyMatrix( m1 );

			return this;

		};

	}(),

	func (geo *BufferGeometry) RotateZ() {

		// rotate geometry around world z-axis

		var m1;

		return function rotateZ( angle ) {

			if ( m1 === undefined ) m1 = new Matrix4();

			m1.makeRotationZ( angle );

			geo.ApplyMatrix( m1 );

			return this;

		};

	}(),

	func (geo *BufferGeometry) Translate() {

		// translate geometry

		var m1;

		return function translate( x, y, z ) {

			if ( m1 === undefined ) m1 = new Matrix4();

			m1.makeTranslation( x, y, z );

			geo.ApplyMatrix( m1 );

			return this;

		};

	}(),

	func (geo *BufferGeometry) Scale() {

		// scale geometry

		var m1;

		return function scale( x, y, z ) {

			if ( m1 === undefined ) m1 = new Matrix4();

			m1.makeScale( x, y, z );

			geo.ApplyMatrix( m1 );

			return this;

		};

	}(),

	func (geo *BufferGeometry) LookAt() {

		var obj;

		return function lookAt( vector ) {

			if ( obj === undefined ) obj = new Object3D();

			obj.lookAt( vector );

			obj.updateMatrix();

			geo.ApplyMatrix( obj.matrix );

		};

	}(),

	func (geo *BufferGeometry) Center() {

		geo.ComputeBoundingBox();

		var offset = geo.BoundingBox.getCenter().negate();

		geo.Translate( offset.x, offset.y, offset.z );

		return offset;

	}

	func (geo *BufferGeometry) SetFromObject( object ) {

		// console.log( "THREE.BufferGeometry.setFromObject(). Converting", object, this );

		var geometry = object.geometry;

		if ( object.isPoints || object.isLine ) {

			var positions = new Float32BufferAttribute( geometry.vertices.length * 3, 3 );
			var colors = new Float32BufferAttribute( geometry.colors.length * 3, 3 );

			geo.AddAttribute( "position", positions.copyVector3sArray( geometry.vertices ) );
			geo.AddAttribute( "color", colors.copyColorsArray( geometry.colors ) );

			if ( geometry.lineDistances && geometry.lineDistances.length === geometry.vertices.length ) {

				var lineDistances = new Float32BufferAttribute( geometry.lineDistances.length, 1 );

				geo.AddAttribute( "lineDistance", lineDistances.copyArray( geometry.lineDistances ) );

			}

			if ( geometry.boundingSphere !== null ) {

				geo.BoundingSphere = geometry.boundingSphere.clone();

			}

			if ( geometry.boundingBox !== null ) {

				geo.BoundingBox = geometry.boundingBox.clone();

			}

		} else if ( object.isMesh ) {

			if ( geometry && geometry.isGeometry ) {

				geo.FromGeometry( geometry );

			}

		}

		return this;

	}

	func (geo *BufferGeometry) UpdateFromObject( object ) {

		var geometry = object.geometry;

		if ( object.isMesh ) {

			var direct = geometry.__directGeometry;

			if ( geometry.elementsNeedUpdate === true ) {

				direct = undefined;
				geometry.elementsNeedUpdate = false;

			}

			if ( direct === undefined ) {

				return geo.FromGeometry( geometry );

			}

			direct.verticesNeedUpdate = geometry.verticesNeedUpdate;
			direct.normalsNeedUpdate = geometry.normalsNeedUpdate;
			direct.colorsNeedUpdate = geometry.colorsNeedUpdate;
			direct.uvsNeedUpdate = geometry.uvsNeedUpdate;
			direct.groupsNeedUpdate = geometry.groupsNeedUpdate;

			geometry.verticesNeedUpdate = false;
			geometry.normalsNeedUpdate = false;
			geometry.colorsNeedUpdate = false;
			geometry.uvsNeedUpdate = false;
			geometry.groupsNeedUpdate = false;

			geometry = direct;

		}

		var attribute;

		if ( geometry.verticesNeedUpdate === true ) {

			attribute = geo.Attributes.position;

			if ( attribute !== undefined ) {

				attribute.copyVector3sArray( geometry.vertices );
				attribute.needsUpdate = true;

			}

			geometry.verticesNeedUpdate = false;

		}

		if ( geometry.normalsNeedUpdate === true ) {

			attribute = geo.Attributes.normal;

			if ( attribute !== undefined ) {

				attribute.copyVector3sArray( geometry.normals );
				attribute.needsUpdate = true;

			}

			geometry.normalsNeedUpdate = false;

		}

		if ( geometry.colorsNeedUpdate === true ) {

			attribute = geo.Attributes.color;

			if ( attribute !== undefined ) {

				attribute.copyColorsArray( geometry.colors );
				attribute.needsUpdate = true;

			}

			geometry.colorsNeedUpdate = false;

		}

		if ( geometry.uvsNeedUpdate ) {

			attribute = geo.Attributes.uv;

			if ( attribute !== undefined ) {

				attribute.copyVector2sArray( geometry.uvs );
				attribute.needsUpdate = true;

			}

			geometry.uvsNeedUpdate = false;

		}

		if ( geometry.lineDistancesNeedUpdate ) {

			attribute = geo.Attributes.lineDistance;

			if ( attribute !== undefined ) {

				attribute.copyArray( geometry.lineDistances );
				attribute.needsUpdate = true;

			}

			geometry.lineDistancesNeedUpdate = false;

		}

		if ( geometry.groupsNeedUpdate ) {

			geometry.computeGroups( object.geometry );
			geo.Groups = geometry.groups;

			geometry.groupsNeedUpdate = false;

		}

		return this;

	}

	func (geo *BufferGeometry) FromGeometry( geometry ) {

		geometry.__directGeometry = new DirectGeometry().fromGeometry( geometry );

		return geo.FromDirectGeometry( geometry.__directGeometry );

	}

	func (geo *BufferGeometry) FromDirectGeometry( geometry ) {

		var positions = new Float32Array( geometry.vertices.length * 3 );
		geo.AddAttribute( "position", new BufferAttribute( positions, 3 ).copyVector3sArray( geometry.vertices ) );

		if ( geometry.normals.length > 0 ) {

			var normals = new Float32Array( geometry.normals.length * 3 );
			geo.AddAttribute( "normal", new BufferAttribute( normals, 3 ).copyVector3sArray( geometry.normals ) );

		}

		if ( geometry.colors.length > 0 ) {

			var colors = new Float32Array( geometry.colors.length * 3 );
			geo.AddAttribute( "color", new BufferAttribute( colors, 3 ).copyColorsArray( geometry.colors ) );

		}

		if ( geometry.uvs.length > 0 ) {

			var uvs = new Float32Array( geometry.uvs.length * 2 );
			geo.AddAttribute( "uv", new BufferAttribute( uvs, 2 ).copyVector2sArray( geometry.uvs ) );

		}

		if ( geometry.uvs2.length > 0 ) {

			var uvs2 = new Float32Array( geometry.uvs2.length * 2 );
			geo.AddAttribute( "uv2", new BufferAttribute( uvs2, 2 ).copyVector2sArray( geometry.uvs2 ) );

		}

		if ( geometry.indices.length > 0 ) {

			var TypeArray = arrayMax( geometry.indices ) > 65535 ? Uint32Array : Uint16Array;
			var indices = new TypeArray( geometry.indices.length * 3 );
			geo.SetIndex( new BufferAttribute( indices, 1 ).copyIndicesArray( geometry.indices ) );

		}

		// groups

		geo.Groups = geometry.groups;

		// morphs

		for ( var name in geometry.morphTargets ) {

			var array = [];
			var morphTargets = geometry.morphTargets[ name ];

			for ( var i = 0, l = morphTargets.length; i < l; i ++ ) {

				var morphTarget = morphTargets[ i ];

				var attribute = new Float32BufferAttribute( morphTarget.length * 3, 3 );

				array.push( attribute.copyVector3sArray( morphTarget ) );

			}

			geo.MorphAttributes[ name ] = array;

		}

		// skinning

		if ( geometry.skinIndices.length > 0 ) {

			var skinIndices = new Float32BufferAttribute( geometry.skinIndices.length * 4, 4 );
			geo.AddAttribute( "skinIndex", skinIndices.copyVector4sArray( geometry.skinIndices ) );

		}

		if ( geometry.skinWeights.length > 0 ) {

			var skinWeights = new Float32BufferAttribute( geometry.skinWeights.length * 4, 4 );
			geo.AddAttribute( "skinWeight", skinWeights.copyVector4sArray( geometry.skinWeights ) );

		}

		//

		if ( geometry.boundingSphere !== null ) {

			geo.BoundingSphere = geometry.boundingSphere.clone();

		}

		if ( geometry.boundingBox !== null ) {

			geo.BoundingBox = geometry.boundingBox.clone();

		}

		return this;

	}

	func (geo *BufferGeometry) ComputeBoundingBox() {

		if ( geo.BoundingBox === null ) {

			geo.BoundingBox = new Box3();

		}

		var position = geo.Attributes.position;

		if ( position !== undefined ) {

			geo.BoundingBox.setFromBufferAttribute( position );

		} else {

			geo.BoundingBox.makeEmpty();

		}

		if ( isNaN( geo.BoundingBox.min.x ) || isNaN( geo.BoundingBox.min.y ) || isNaN( geo.BoundingBox.min.z ) ) {

			console.error( "THREE.BufferGeometry.computeBoundingBox: Computed min/max have NaN values. The "position" attribute is likely to have NaN values.", this );

		}

	}

	func (geo *BufferGeometry) ComputeBoundingSphere() {

		var box = new Box3();
		var vector = new Vector3();

		return function computeBoundingSphere() {

			if ( geo.BoundingSphere === null ) {

				geo.BoundingSphere = new Sphere();

			}

			var position = geo.Attributes.position;

			if ( position ) {

				var center = geo.BoundingSphere.center;

				box.setFromBufferAttribute( position );
				box.getCenter( center );

				// hoping to find a boundingSphere with a radius smaller than the
				// boundingSphere of the boundingBox: sqrt(3) smaller in the best case

				var maxRadiusSq = 0;

				for ( var i = 0, il = position.count; i < il; i ++ ) {

					vector.x = position.getX( i );
					vector.y = position.getY( i );
					vector.z = position.getZ( i );
					maxRadiusSq = Math.max( maxRadiusSq, center.distanceToSquared( vector ) );

				}

				geo.BoundingSphere.radius = Math.sqrt( maxRadiusSq );

				if ( isNaN( geo.BoundingSphere.radius ) ) {

					console.error( "THREE.BufferGeometry.computeBoundingSphere(): Computed radius is NaN. The "position" attribute is likely to have NaN values.", this );

				}

			}

		};

	}(),

	func (geo *BufferGeometry) ComputeFaceNormals() {

		// backwards compatibility

	}

	func (geo *BufferGeometry) ComputeVertexNormals() {

		var index = geo.Index;
		var attributes = geo.Attributes;
		var groups = geo.Groups;

		if ( attributes.position ) {

			var positions = attributes.position.array;

			if ( attributes.normal === undefined ) {

				geo.AddAttribute( "normal", new BufferAttribute( new Float32Array( positions.length ), 3 ) );

			} else {

				// reset existing normals to zero

				var array = attributes.normal.array;

				for ( var i = 0, il = array.length; i < il; i ++ ) {

					array[ i ] = 0;

				}

			}

			var normals = attributes.normal.array;

			var vA, vB, vC;
			var pA = new Vector3(), pB = new Vector3(), pC = new Vector3();
			var cb = new Vector3(), ab = new Vector3();

			// indexed elements

			if ( index ) {

				var indices = index.array;

				if ( groups.length === 0 ) {

					geo.AddGroup( 0, indices.length );

				}

				for ( var j = 0, jl = groups.length; j < jl; ++ j ) {

					var group = groups[ j ];

					var start = group.start;
					var count = group.count;

					for ( var i = start, il = start + count; i < il; i += 3 ) {

						vA = indices[ i + 0 ] * 3;
						vB = indices[ i + 1 ] * 3;
						vC = indices[ i + 2 ] * 3;

						pA.fromArray( positions, vA );
						pB.fromArray( positions, vB );
						pC.fromArray( positions, vC );

						cb.subVectors( pC, pB );
						ab.subVectors( pA, pB );
						cb.cross( ab );

						normals[ vA ] += cb.x;
						normals[ vA + 1 ] += cb.y;
						normals[ vA + 2 ] += cb.z;

						normals[ vB ] += cb.x;
						normals[ vB + 1 ] += cb.y;
						normals[ vB + 2 ] += cb.z;

						normals[ vC ] += cb.x;
						normals[ vC + 1 ] += cb.y;
						normals[ vC + 2 ] += cb.z;

					}

				}

			} else {

				// non-indexed elements (unconnected triangle soup)

				for ( var i = 0, il = positions.length; i < il; i += 9 ) {

					pA.fromArray( positions, i );
					pB.fromArray( positions, i + 3 );
					pC.fromArray( positions, i + 6 );

					cb.subVectors( pC, pB );
					ab.subVectors( pA, pB );
					cb.cross( ab );

					normals[ i ] = cb.x;
					normals[ i + 1 ] = cb.y;
					normals[ i + 2 ] = cb.z;

					normals[ i + 3 ] = cb.x;
					normals[ i + 4 ] = cb.y;
					normals[ i + 5 ] = cb.z;

					normals[ i + 6 ] = cb.x;
					normals[ i + 7 ] = cb.y;
					normals[ i + 8 ] = cb.z;

				}

			}

			geo.NormalizeNormals();

			attributes.normal.needsUpdate = true;

		}

	}

	func (geo *BufferGeometry) Merge( geometry, offset ) {

		if ( ( geometry && geometry.isBufferGeometry ) === false ) {

			console.error( "THREE.BufferGeometry.merge(): geometry not an instance of THREE.BufferGeometry.", geometry );
			return;

		}

		if ( offset === undefined ) offset = 0;

		var attributes = geo.Attributes;

		for ( var key in attributes ) {

			if ( geometry.attributes[ key ] === undefined ) continue;

			var attribute1 = attributes[ key ];
			var attributeArray1 = attribute1.array;

			var attribute2 = geometry.attributes[ key ];
			var attributeArray2 = attribute2.array;

			var attributeSize = attribute2.itemSize;

			for ( var i = 0, j = attributeSize * offset; i < attributeArray2.length; i ++, j ++ ) {

				attributeArray1[ j ] = attributeArray2[ i ];

			}

		}

		return this;

	}

	func (geo *BufferGeometry) NormalizeNormals() {

		var normals = geo.Attributes.normal;

		var x, y, z, n;

		for ( var i = 0, il = normals.count; i < il; i ++ ) {

			x = normals.getX( i );
			y = normals.getY( i );
			z = normals.getZ( i );

			n = 1.0 / Math.sqrt( x * x + y * y + z * z );

			normals.setXYZ( i, x * n, y * n, z * n );

		}

	}

	func (geo *BufferGeometry) ToNonIndexed() {

		if ( geo.Index === null ) {

			console.warn( "THREE.BufferGeometry.toNonIndexed(): Geometry is already non-indexed." );
			return this;

		}

		var geometry2 = new BufferGeometry();

		var indices = geo.Index.array;
		var attributes = geo.Attributes;

		for ( var name in attributes ) {

			var attribute = attributes[ name ];

			var array = attribute.array;
			var itemSize = attribute.itemSize;

			var array2 = new array.constructor( indices.length * itemSize );

			var index = 0, index2 = 0;

			for ( var i = 0, l = indices.length; i < l; i ++ ) {

				index = indices[ i ] * itemSize;

				for ( var j = 0; j < itemSize; j ++ ) {

					array2[ index2 ++ ] = array[ index ++ ];

				}

			}

			geometry2.addAttribute( name, new BufferAttribute( array2, itemSize ) );

		}

		return geometry2;

	}

	func (geo *BufferGeometry) ToJSON() {

		var data = {
			metadata: {
				version: 4.5,
				type: "BufferGeometry",
				generator: "BufferGeometry.toJSON"
			}
		};

		// standard BufferGeometry serialization

		data.uuid = geo.Uuid;
		data.type = geo.Type;
		if ( geo.Name !== "" ) data.name = geo.Name;

		if ( geo.Parameters !== undefined ) {

			var parameters = geo.Parameters;

			for ( var key in parameters ) {

				if ( parameters[ key ] !== undefined ) data[ key ] = parameters[ key ];

			}

			return data;

		}

		data.data = { attributes: {} };

		var index = geo.Index;

		if ( index !== null ) {

			var array = Array.prototype.slice.call( index.array );

			data.data.index = {
				type: index.array.constructor.name,
				array: array
			};

		}

		var attributes = geo.Attributes;

		for ( var key in attributes ) {

			var attribute = attributes[ key ];

			var array = Array.prototype.slice.call( attribute.array );

			data.data.attributes[ key ] = {
				itemSize: attribute.itemSize,
				type: attribute.array.constructor.name,
				array: array,
				normalized: attribute.normalized
			};

		}

		var groups = geo.Groups;

		if ( groups.length > 0 ) {

			data.data.groups = JSON.parse( JSON.stringify( groups ) );

		}

		var boundingSphere = geo.BoundingSphere;

		if ( boundingSphere !== null ) {

			data.data.boundingSphere = {
				center: boundingSphere.center.toArray(),
				radius: boundingSphere.radius
			};

		}

		return data;

	}*/

//func (geo *BufferGeometry) Clone() {

/*
 // Handle primitives

 var parameters = geo.Parameters;

 if ( parameters !== undefined ) {

 var values = [];

 for ( var key in parameters ) {

 values.push( parameters[ key ] );

 }

 var geometry = Object.create( geo.Constructor.prototype );
 geo.Constructor.apply( geometry, values );
 return geometry;

 }

 return new geo.Constructor().copy( this );
*/

//return new BufferGeometry().copy( this );

//}

/*func (geo *BufferGeometry) Copy( source ) {

		var name, i, l;

		// reset

		geo.Index = null;
		geo.Attributes = {};
		geo.MorphAttributes = {};
		geo.Groups = [];
		geo.BoundingBox = null;
		geo.BoundingSphere = null;

		// name

		geo.Name = source.name;

		// index

		var index = source.index;

		if ( index !== null ) {

			geo.SetIndex( index.clone() );

		}

		// attributes

		var attributes = source.attributes;

		for ( name in attributes ) {

			var attribute = attributes[ name ];
			geo.AddAttribute( name, attribute.clone() );

		}

		// morph attributes

		var morphAttributes = source.morphAttributes;

		for ( name in morphAttributes ) {

			var array = [];
			var morphAttribute = morphAttributes[ name ]; // morphAttribute: array of Float32BufferAttributes

			for ( i = 0, l = morphAttribute.length; i < l; i ++ ) {

				array.push( morphAttribute[ i ].clone() );

			}

			geo.MorphAttributes[ name ] = array;

		}

		// groups

		var groups = source.groups;

		for ( i = 0, l = groups.length; i < l; i ++ ) {

			var group = groups[ i ];
			geo.AddGroup( group.start, group.count, group.materialIndex );

		}

		// bounding box

		var boundingBox = source.boundingBox;

		if ( boundingBox !== null ) {

			geo.BoundingBox = boundingBox.clone();

		}

		// bounding sphere

		var boundingSphere = source.boundingSphere;

		if ( boundingSphere !== null ) {

			geo.BoundingSphere = boundingSphere.clone();

		}

		// draw range

		geo.DrawRange.start = source.drawRange.start;
		geo.DrawRange.count = source.drawRange.count;

		return this;

	}

	func (geo *BufferGeometry) Dispose() {

		geo.DispatchEvent( { type: "dispose" } );

	}

} );


export { BufferGeometry };
*/
